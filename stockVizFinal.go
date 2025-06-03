package main

import "fmt"

const stockLength int = 30
const maxPorto int = 6

type Stockstruct struct {
	Name   string
	Prices [6]float64
}

type Portfoliostruct struct {
	Name          [maxPorto]string
	BuyPrice      [maxPorto]float64
	Quantity      [maxPorto]int 
	Profit        [maxPorto]float64
	ReturnPercent [maxPorto]float64
}

type Stock [stockLength]Stockstruct

type Portfolio Portfoliostruct

func main() {
	var choice int
	var balance float64
	var selectionMethod, searchMethod int = 1, 1 // Default to Selection Sort for display, Sequential Search

	welcomeScreen()
	fmt.Scan(&choice)

	for choice != 5 {
		switch choice {
		case 1:
			idrToUsdConverter(&balance)
		case 2:
			balanceModify(&balance)
		case 3:
			stockProfitLossViz(&balance, selectionMethod, searchMethod)
		case 4:
			options(&selectionMethod, &searchMethod)
		default:
			fmt.Printf("Invalid Choice!\n")
		}
		fmt.Printf("\n============================================================\n")
		welcomeScreen()
		fmt.Scan(&choice)
	}
	fmt.Printf("Closing....... Thank You\n")
}

func welcomeScreen() {
	/*
	   I.S = None
	   F.S = Shows the user the title text and option list
	*/
	var ascii string
	ascii = `
#***********************************************************
#***********************************************************
#***********************************************########****
#************************************************######*****
#***********************************************#######*****
#*********************************************####**##******
#*******************************************####************
#*****************************************####**************
#***************************************###*****************
#*************************************###*******************
#**********************###**********###*********************
#*********************########****###***********************
#******##************##*****#######*************************
#****#####*********###**********#***************************
#***###**###******###***************************************
#*###******###***###****************************************
####*********######*****************************************
##*************##*******************************************
#***********************************************************
############################################################
============================================================
Welcome to Stock Tracker (US ONLY)
============================================================
1. Indonesian Rupiah To US Dollar Conversion
2. Check / Increase / Decrease the balance in USD
3. Stock Vizualization
4. Options
5. Exit
------------------------------------------------------------
Your Option: `
	fmt.Printf("%s", ascii)
}

func idrToUsdConverter(bal *float64) {
	/*
	   I.S = None
	   F.S = Prints the converted value to user and optionally saves the value
	*/
	var convertChoice int
	var userInput float64
	var saveVal string
	var rate float64 = 16474.00 // Standard rate (example: 1 USD = 16474 IDR)

	fmt.Printf("\nChoose Conversion Method:\n1. Use default rate (1 USD = %.2f IDR)\n2. Custom Rate\nYour Option: ", rate)
	fmt.Scan(&convertChoice)

	validChoice := false
	for !validChoice {
		if convertChoice == 1 || convertChoice == 2 {
			validChoice = true
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options (1 or 2)\nYour Option: ")
			fmt.Scan(&convertChoice)
		}
	}

	if convertChoice == 2 {
		fmt.Printf("Set Your Prefered Rate (IDR per 1 USD): ")
		fmt.Scan(&rate)
		fmt.Printf("\n")
	}

	fmt.Printf("How Much Would You Convert: Rp.")
	fmt.Scan(&userInput)
	convertedAmount := userInput / rate
	fmt.Printf("\nYour Conversion Result Is: $%.2f\n", convertedAmount)

	fmt.Printf("\nDo you wish to save the conversion value to the balance? (Yes/No)\nYour Option: ")
	fmt.Scan(&saveVal)
	if saveVal == "Yes" || saveVal == "yes" || saveVal == "y" || saveVal == "Y" {
		*bal += convertedAmount
		fmt.Printf("\nValue Saved. New Balance: $%.2f\n", *bal)
	}
	fmt.Printf("Returning to the main screen.......\n")
}

func balanceModify(bal *float64) {
	/*
	   I.S = None or will be recieve the user balance
	   F.S = Prints the modified balance to user
	*/
	var userChoice int
	var userVal float64

	fmt.Printf("\nYour Balance: $%.2f\n", *bal)
	fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: ")
	fmt.Scan(&userChoice)

	for userChoice != 2 {
		if userChoice == 1 {
			fmt.Printf("\nEnter a value to add or subtract (use '-' for reduction):\nYour Option: ")
			fmt.Scan(&userVal)
			*bal += userVal
			fmt.Printf("\nYour new Balance: $%.2f\n", *bal)
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options (1 or 2)\n")
		}
		fmt.Printf("\nPlease choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: ")
		fmt.Scan(&userChoice)
	}
	fmt.Printf("Returning to the main screen.......\n")
}

func sortStocksByName(stocksToSort *Stock) {
	var i, j, minIdx int
	var temp Stockstruct
	var n int = stockLength

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if stocksToSort[j].Name < stocksToSort[minIdx].Name {
				minIdx = j
			}
		}
		temp = stocksToSort[minIdx]
		stocksToSort[minIdx] = stocksToSort[i]
		stocksToSort[i] = temp
	}
}

func stockProfitLossViz(bal *float64, displaySortMethod int, searchMethod int) {
	/*
	   I.S = User balance, chosen display sort method, chosen search method
	   F.S = Shows the user the list of stocks and allows trading simulation
	*/
	var stocksData Stock = Stock{
		{"AAPL", [6]float64{172.34, 168.21, 175.89, 170.56, 173.12, 174.50}},
		{"MSFT", [6]float64{324.56, 329.12, 327.45, 330.78, 328.90, 331.10}},
		{"GOOGL", [6]float64{170.45, 172.67, 171.12, 173.34, 170.78, 174.00}},
		{"AMZN", [6]float64{134.22, 137.45, 136.78, 138.90, 135.67, 139.20}},
		{"NVDA", [6]float64{690.12, 700.45, 695.78, 702.34, 698.56, 703.90}},
		{"TSLA", [6]float64{210.34, 205.12, 215.67, 212.45, 213.89, 216.00}},
		{"META", [6]float64{312.67, 315.12, 310.45, 317.89, 313.56, 318.20}},
		{"BRK.B", [6]float64{340.12, 342.45, 341.78, 343.90, 340.67, 344.10}},
		{"JPM", [6]float64{158.45, 160.12, 159.67, 161.34, 158.90, 162.00}},
		{"V", [6]float64{232.12, 234.45, 233.78, 235.90, 232.67, 236.10}},
		{"JNJ", [6]float64{168.90, 170.45, 169.78, 171.12, 170.34, 171.80}},
		{"UNH", [6]float64{510.34, 512.12, 511.67, 513.45, 510.89, 514.00}},
		{"XOM", [6]float64{104.56, 106.12, 105.78, 107.45, 104.90, 107.80}},
		{"PG", [6]float64{152.34, 154.12, 153.67, 155.45, 152.89, 155.90}},
		{"HD", [6]float64{320.45, 322.12, 321.67, 323.34, 320.90, 324.00}},
		{"MA", [6]float64{390.12, 392.45, 391.78, 393.90, 390.67, 394.20}},
		{"CVX", [6]float64{162.78, 164.12, 163.67, 165.45, 162.89, 165.90}},
		{"LLY", [6]float64{780.34, 782.12, 781.67, 783.45, 780.89, 784.00}},
		{"MRK", [6]float64{112.45, 114.12, 113.67, 115.34, 112.90, 115.80}},
		{"PEP", [6]float64{182.34, 184.12, 183.67, 185.45, 182.89, 185.90}},
		{"COST", [6]float64{710.78, 712.12, 711.67, 713.45, 710.89, 714.20}},
		{"AVGO", [6]float64{900.45, 902.12, 901.67, 903.34, 900.90, 904.00}},
		{"ADBE", [6]float64{570.12, 572.45, 571.78, 573.90, 570.67, 574.20}},
		{"BAC", [6]float64{28.45, 29.12, 28.67, 29.34, 28.90, 29.50}},
		{"ABBV", [6]float64{162.12, 164.45, 163.78, 165.90, 162.67, 166.10}},
		{"DIS", [6]float64{102.34, 104.12, 103.67, 105.45, 102.89, 105.90}},
		{"KO", [6]float64{61.12, 62.45, 61.78, 63.90, 61.67, 64.10}},
		{"MCD", [6]float64{288.34, 290.12, 289.67, 291.45, 288.89, 292.00}},
		{"NFLX", [6]float64{420.45, 422.12, 421.67, 423.34, 420.90, 424.00}},
		{"WMT", [6]float64{162.12, 164.45, 163.78, 165.90, 162.67, 166.10}},
	}
	var userChoice, i int
	var userPortfolio Portfolio
	var exitMenu bool = false

	if searchMethod == 2 { // Binary Search
		sortStocksByName(&stocksData)
	}

	fmt.Printf("\nPlease choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\nYour Option: ")
	fmt.Scan(&userChoice)

	for !exitMenu {
		if userChoice == 1 {
			fmt.Printf("\nAvailable Stocks (Ticker Symbols):\n")
			for i = 0; i < stockLength; i += 5 {
				fmt.Printf("%s %s %s %s %s\n", stocksData[i].Name, stocksData[i+1].Name, stocksData[i+2].Name, stocksData[i+3].Name, stocksData[i+4].Name)
			}
		} else if userChoice == 2 {
			investSellStocks(&stocksData, &userPortfolio, bal, displaySortMethod, searchMethod)
		} else if userChoice == 3 {
			exitMenu = true
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options (1-3)\n")
		}

		if !exitMenu {
			fmt.Printf("\nPlease choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\nYour Option: ")
			fmt.Scan(&userChoice)
		}
	}
	fmt.Printf("Returning to the main screen.......\n")
}

func investSellStocks(stocks *Stock, userPortfolio *Portfolio, bal *float64, displaySortMethod int, searchMethod int) {
	/*
		I.S = Stocks data, User Portfolio, User Balance, display sort method, search method
		F.S = Simulates trading days, allowing user to buy/sell stocks
	*/
	var stockOptionChoice int
	var portfolioStockCount int = 0
	var currentDay int = 1
	var exitTradingDayLoop bool

	// Trading simulation for 5 days
	for currentDay <= 5 {
		fmt.Printf("\n--- Trading Day %d ---\n", currentDay)
		fmt.Printf("Your Balance: $%.2f\n", *bal)
		fmt.Printf("Your Portfolio:\n")
		showPortoResult(stocks, userPortfolio, currentDay, searchMethod)

		exitTradingDayLoop = false
		for !exitTradingDayLoop {
			fmt.Print("\nSelect the following options for Day ", currentDay, ":\n")
			fmt.Print("1. View previous day's highest gainers\n")
			fmt.Print("2. Buy a stock\n")
			fmt.Print("3. Sell a stock\n")
			fmt.Print("4. Skip to next day\n")
			fmt.Print("5. Exit trading simulation\n")
			fmt.Print("Your Option: ")
			fmt.Scan(&stockOptionChoice)

			if stockOptionChoice == 1 {
				priceIdxForSort := currentDay - 1
				if priceIdxForSort < 0 {
					priceIdxForSort = 0
				}
				if displaySortMethod == 1 { // Selection Sort
					sortHighestPriceSelection(*stocks, priceIdxForSort)
				} else { // Insertion Sort
					sortHighestPriceInsertion(*stocks, priceIdxForSort)
				}
			} else if stockOptionChoice == 2 {
				buyStock(stocks, userPortfolio, bal, currentDay, searchMethod, &portfolioStockCount)
			} else if stockOptionChoice == 3 {
				sellStock(stocks, userPortfolio, bal, currentDay, searchMethod, &portfolioStockCount)
			} else if stockOptionChoice == 4 {
				fmt.Println("Skipping to the next day...")
				exitTradingDayLoop = true
			} else if stockOptionChoice == 5 {
				fmt.Println("Exiting trading simulation...")
				return
			} else {
				fmt.Printf("Invalid Value... Please Select The Correct Options (1-5)\n")
			}
		}
		currentDay++
	}
	fmt.Println("\n--- End of 5 Trading Days ---")
}

func sortHighestPriceSelection(stocksData Stock, priceDayIndex int) {
	/*
	   I.S = Stocks data (passed by value, so it's a copy) and priceDayIndex (0 for Day 0 prices, etc.)
	   F.S = Sorts and displays the top 5 stocks by the highest price for the given priceDayIndex
	*/
	var i, min, pass int
	var temp Stockstruct
	var n int = stockLength

	if priceDayIndex < 0 || priceDayIndex >= 6 {
		fmt.Printf("Invalid priceDayIndex for sorting: %d\n", priceDayIndex)
		return
	}

	for pass = 0; pass < n-1; pass++ {
		min = pass
		for i = pass + 1; i < n; i++ {
			if stocksData[min].Prices[priceDayIndex] < stocksData[i].Prices[priceDayIndex] {
				min = i
			}
		}
		temp = stocksData[pass]
		stocksData[pass] = stocksData[min]
		stocksData[min] = temp
	}

	fmt.Printf("\nTop 5 highest prices from Day %d:\n", priceDayIndex)
	displayCount := 5
	if n < 5 {
		displayCount = n
	}
	for i = 0; i < displayCount; i++ {
		fmt.Printf("%d. %s: $%.2f\n", i+1, stocksData[i].Name, stocksData[i].Prices[priceDayIndex])
	}
}

func sortHighestPriceInsertion(stocksData Stock, priceDayIndex int) {
	/*
	   I.S = Stocks data (passed by value, so it's a copy) and priceDayIndex
	   F.S = Sorts and displays the top 5 stocks by the highest price for the given priceDayIndex
	*/
	var pass, j, i int
	var temp Stockstruct
	var n int = stockLength

	if priceDayIndex < 0 || priceDayIndex >= 6 {
		fmt.Printf("Invalid priceDayIndex for sorting: %d\n", priceDayIndex)
		return
	}

	for pass = 1; pass < n; pass++ {
		temp = stocksData[pass]
		j = pass
		for j > 0 && stocksData[j-1].Prices[priceDayIndex] < temp.Prices[priceDayIndex] {
			stocksData[j] = stocksData[j-1]
			j--
		}
		stocksData[j] = temp
	}

	fmt.Printf("\nTop 5 highest prices from Day %d:\n", priceDayIndex)
	displayCount := 5
	if n < 5 {
		displayCount = n
	}
	for i = 0; i < displayCount; i++ {
		fmt.Printf("%d. %s: $%.2f\n", i+1, stocksData[i].Name, stocksData[i].Prices[priceDayIndex])
	}
}

func searchUsingBinSearch(stocksToSearch Stock, stockName string) int {
	/*
	   I.S = Stocks data (assumed to be sorted by Name)
	   F.S = Returns the index of the stock if found or -1 if not found
	*/
	var low int = 0
	var high int = stockLength - 1
	var found bool = false
	var mid int = -1

	for low <= high && !found {
		mid = (low + high) / 2
		if stocksToSearch[mid].Name == stockName {
			found = true
		} else if stocksToSearch[mid].Name < stockName {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		return mid
	}
	return -1
}

func searchUsingSeqSearch(stocksToSearch Stock, stockName string) int {
	/*
	   I.S = Stocks data
	   F.S = Returns the index of the stock if found or -1 if not found
	*/
	var i int
	var found bool = false
	var foundIdx int = -1

	for i = 0; i < stockLength && !found; i++ {
		if stocksToSearch[i].Name == stockName {
			foundIdx = i
			found = true
		}
	}
	return foundIdx
}

func options(selectionMethod, searchMethod *int) {
	/*
		I.S = Pointers to selectionMethod and searchMethod
		F.S = User sets their preferred sorting and searching methods
	*/
	exitOptions := false
	var userChoice int

	for !exitOptions {
		fmt.Printf("\nOptions Menu:\n")
		fmt.Printf("Current Sorting Method for display: ")
		if *selectionMethod == 1 {
			fmt.Printf("Selection Sort\n")
		} else {
			fmt.Printf("Insertion Sort\n")
		}
		fmt.Printf("Current Search Method for trading: ")
		if *searchMethod == 1 {
			fmt.Printf("Sequential Search\n")
		} else {
			fmt.Printf("Binary Search (Requires stocks to be sorted by name)\n")
		}

		fmt.Printf("\n1. Change Sorting Method (for display)\n2. Change Search Method (for trading)\n3. Exit Options\nYour Option: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Printf("Select the sorting method (default: Selection Sort):\n1. Selection Sort\n2. Insertion Sort\nYour Option: ")
			fmt.Scan(selectionMethod)
			if *selectionMethod != 1 && *selectionMethod != 2 {
				fmt.Println("Invalid selection, defaulting to Selection Sort.")
				*selectionMethod = 1
			}
		} else if userChoice == 2 {
			fmt.Printf("Select the search method (default: Sequential Search):\n1. Sequential Search\n2. Binary Search\nYour Option: ")
			fmt.Scan(searchMethod)
			if *searchMethod != 1 && *searchMethod != 2 {
				fmt.Println("Invalid selection, defaulting to Sequential Search.")
				*searchMethod = 1
			}
			if *searchMethod == 2 {
				fmt.Println("Note: Binary search will sort the stock list by name at the start of trading simulation.")
			}
		} else if userChoice == 3 {
			exitOptions = true
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options (1-3)\n")
		}
	}
	fmt.Printf("Returning to the main screen.......\n")
}

func showPortoResult(stocks *Stock, userPortfolio *Portfolio, currentDay int, searchMethod int) {
	/*
	   I.S = Stocks data, user's portfolio, current trading day
	   F.S = Displays the current state of the user's portfolio
	*/
	var hasHoldings bool = false
	var heldStockName string
	var actualStockIndex int
	var currentPrice, profit, initialInvestment, returnPercent float64

	fmt.Println("--- Your Portfolio Status ---")
	for i := 0; i < maxPorto; i++ {
		if userPortfolio.Quantity[i] > 0 {
			hasHoldings = true
			heldStockName = userPortfolio.Name[i]

			if searchMethod == 2 {
				actualStockIndex = searchUsingBinSearch(*stocks, heldStockName)
			} else {
				actualStockIndex = searchUsingSeqSearch(*stocks, heldStockName)
			}

			if actualStockIndex != -1 {
				currentPrice = stocks[actualStockIndex].Prices[currentDay]

				fmt.Printf("Stock: %s | Qty: %d | Buy Price: $%.2f ", userPortfolio.Name[i], userPortfolio.Quantity[i], userPortfolio.BuyPrice[i])

				if currentDay > 1 {
					profit = (currentPrice - userPortfolio.BuyPrice[i]) * float64(userPortfolio.Quantity[i])
					userPortfolio.Profit[i] = profit

					initialInvestment = userPortfolio.BuyPrice[i] * float64(userPortfolio.Quantity[i])
					returnPercent = (profit / initialInvestment) * 100

					userPortfolio.ReturnPercent[i] = returnPercent

					fmt.Printf("| Current Price: $%.2f | Profit: $%.2f | Return: %.2f%%\n", currentPrice, userPortfolio.Profit[i], userPortfolio.ReturnPercent[i])
				} else {
					fmt.Printf("| Current Price: $%.2f (Profit not calculated yet for Day 0)\n", currentPrice)
				}
			} else {
				fmt.Printf("Error: Stock %s from portfolio not found in stock list.\n", heldStockName)
			}
		}
	}
	if !hasHoldings {
		fmt.Println("Your portfolio is currently empty.")
	}
	fmt.Println("-----------------------------")
}

func buyStock(stocks *Stock, userPortfolio *Portfolio, bal *float64, currentDay int, searchMethod int, portfolioStockCount *int) {
	var stockSearch string
	var stockIdx int
	var qty int
	var priceTransactionIndex int
	var transactionPrice, cost float64

	fmt.Printf("\nSearch a stock to buy (enter ticker symbol):\nYour Option: ")
	fmt.Scan(&stockSearch)

	if searchMethod == 1 {
		stockIdx = searchUsingSeqSearch(*stocks, stockSearch)
	} else {
		stockIdx = searchUsingBinSearch(*stocks, stockSearch)
	}

	if stockIdx != -1 {
		priceTransactionIndex = currentDay - 1
		transactionPrice = (*stocks)[stockIdx].Prices[priceTransactionIndex]

		fmt.Printf("Stock Found: %s\n", (*stocks)[stockIdx].Name)
		fmt.Printf("Current Price for buying on Day %d: $%.2f\n", currentDay, transactionPrice)
		fmt.Printf("Your Balance: $%.2f\n", *bal)
		fmt.Printf("How many shares would you like to buy?\nYour Option: ")
		fmt.Scan(&qty)

		if qty <= 0 {
			fmt.Println("Quantity must be positive.")
			return
		}

		cost = transactionPrice * float64(qty)
		if cost <= *bal {
			targetPortfolioSlot := -1
			isExistingHolding := false

			// Check if stock already in portfolio
			for i := 0; i < maxPorto && targetPortfolioSlot == -1; i++ {
				if userPortfolio.Name[i] == (*stocks)[stockIdx].Name {
					targetPortfolioSlot = i
					isExistingHolding = true
				}
			}

			if !isExistingHolding {
				// If it's a new stock, check if portfolio has space for a new distinct stock
				if *portfolioStockCount >= maxPorto {
					fmt.Printf("Portfolio limit reached for distinct stocks! You cannot invest in more than %d different stocks.\n", maxPorto)
					return
				}
				// Find an empty slot for the new stock
				for i := 0; i < maxPorto && targetPortfolioSlot == -1; i++ {
					if userPortfolio.Quantity[i] == 0 {
						targetPortfolioSlot = i
					}
				}
			}

			if targetPortfolioSlot == -1 {
				fmt.Printf("Error: Portfolio might be full.\n")
				return
			}

			*bal -= cost
			if !isExistingHolding { // New stock entry in portfolio
				*portfolioStockCount += 1
				userPortfolio.Name[targetPortfolioSlot] = (*stocks)[stockIdx].Name
				userPortfolio.BuyPrice[targetPortfolioSlot] = transactionPrice
				userPortfolio.Quantity[targetPortfolioSlot] = qty
			} else { // Adding to existing holding
				totalOldCost := userPortfolio.BuyPrice[targetPortfolioSlot] * float64(userPortfolio.Quantity[targetPortfolioSlot])
				totalNewSharesCost := transactionPrice * float64(qty)

				userPortfolio.Quantity[targetPortfolioSlot] += qty
				userPortfolio.BuyPrice[targetPortfolioSlot] = (totalOldCost + totalNewSharesCost) / float64(userPortfolio.Quantity[targetPortfolioSlot]) // Average buy price
			}
			fmt.Printf("Purchase successful! %d shares of %s bought.\n", qty, (*stocks)[stockIdx].Name)
			fmt.Printf("New balance: $%.2f\n", *bal)

		} else {
			fmt.Printf("Insufficient funds. Cost: $%.2f, Your Balance: $%.2f\n", cost, *bal)
		}
	} else {
		fmt.Printf("Stock Ticker '%s' Not Found in the market list.\n", stockSearch)
	}
}

func sellStock(stocks *Stock, userPortfolio *Portfolio, bal *float64, currentDay int, searchMethod int, portfolioStockCount *int) {
	var stockSearch string
	var stockMarketIdx, portfolioIdx int
	var qty int

	if *portfolioStockCount == 0 {
		fmt.Println("Your portfolio is empty. Nothing to sell.")
		return
	}

	fmt.Printf("\nSearch a stock to sell from your portfolio (enter ticker symbol):\nYour Option: ")
	fmt.Scan(&stockSearch)

	// Find the stock in the user's portfolio first
	portfolioIdx = -1
	foundInPortfolio := false
	for i := 0; i < maxPorto && !foundInPortfolio; i++ {
		if userPortfolio.Name[i] == stockSearch && userPortfolio.Quantity[i] > 0 {
			portfolioIdx = i
			foundInPortfolio = true
		}
	}

	if foundInPortfolio {
		// Now find this stock in the main market list to get its current price
		// This step is somewhat redundant if we trust portfolio name, but good for getting market data struct
		if searchMethod == 1 {
			stockMarketIdx = searchUsingSeqSearch(*stocks, userPortfolio.Name[portfolioIdx])
		} else {
			stockMarketIdx = searchUsingBinSearch(*stocks, userPortfolio.Name[portfolioIdx])
		}

		if stockMarketIdx == -1 {
			fmt.Printf("Error: Stock %s found in portfolio but not in market data. This should not happen.\n", userPortfolio.Name[portfolioIdx])
			return
		}

		priceTransactionIndex := currentDay - 1
		if priceTransactionIndex < 0 || priceTransactionIndex >= 6 {
			fmt.Printf("Error: Invalid price index for transaction for stock %s.\n", (*stocks)[stockMarketIdx].Name)
			return
		}
		transactionPrice := (*stocks)[stockMarketIdx].Prices[priceTransactionIndex]

		fmt.Printf("Stock Found in portfolio: %s\n", userPortfolio.Name[portfolioIdx])
		fmt.Printf("Current Price for selling on Day %d: $%.2f\n", currentDay, transactionPrice)
		fmt.Printf("You own %d shares. How many shares would you like to sell?\nYour Option: ", userPortfolio.Quantity[portfolioIdx])
		fmt.Scan(&qty)

		if qty > 0 && qty <= userPortfolio.Quantity[portfolioIdx] {
			proceeds := transactionPrice * float64(qty)
			*bal += proceeds
			userPortfolio.Quantity[portfolioIdx] -= qty

			fmt.Printf("Sale successful! %d shares of %s sold.\n", qty, userPortfolio.Name[portfolioIdx])
			fmt.Printf("New balance: $%.2f\n", *bal)

			if userPortfolio.Quantity[portfolioIdx] == 0 {
				fmt.Printf("You have sold all shares of %s.\n", userPortfolio.Name[portfolioIdx])
				*portfolioStockCount -= 1
				userPortfolio.Name[portfolioIdx] = ""
				userPortfolio.BuyPrice[portfolioIdx] = 0
				userPortfolio.Profit[portfolioIdx] = 0
				userPortfolio.ReturnPercent[portfolioIdx] = 0
			}
		} else if qty <= 0 {
			fmt.Printf("Quantity to sell must be positive.\n")
		} else {
			fmt.Printf("Invalid quantity. You only own %d shares of %s.\n", userPortfolio.Quantity[portfolioIdx], userPortfolio.Name[portfolioIdx])
		}
	} else {
		fmt.Printf("Stock Ticker '%s' Not Found in your portfolio or you have no shares of it.\n", stockSearch)
	}
}

