package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stock struct {
	Name   string
	Prices [6]float64
}

type Portfolio struct {
	Name          [30]string
	buyPrice      [30]float64
	Quantity      [30]int
	Profit        [30]float64
	ReturnPercent [30]float64
}

func main() {
	// Variable Declaration
	var choice int
	var balance float64

	//Random number generator
	rand.Seed(time.Now().UnixNano())

	welcomeScreen()
	fmt.Print("Your Option: ")
	fmt.Scan(&choice)

	// Redirect the user option to the corresponding function and repeats the title screen until the user option is 4 (exit)
	for choice != 4 {
		if choice == 1 {
			idrToUsdConverter(&balance)
		} else if choice == 2 {
			balanceModify(&balance)
		} else if choice == 3 {
			stockProfitLossViz(&balance)
		} else if choice < 1 || choice > 4 {
			fmt.Printf("Invalid Choice! Please enter a number between 1 and 4.\n")
		}
		welcomeScreen()
		fmt.Print("Your Option: ")
		fmt.Scan(&choice)
	}
	// Text to show the program has ended
	fmt.Printf("Closing...... Thank You\n")
}

func welcomeScreen() {
	/*
	   I.S = None
	   F.S = Shows the user the title text and option list
	*/

	// Variable declaration
	var ascii string

	// Initialization of the title screen art and the available options
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
Welcome to US Stock Tracker 
============================================================
1. Indonesian Rupiah To US Dollar Conversion
2. Check / Adjust the balance in USD
3. View / Invest / Sell Stock
4. Exit
------------------------------------------------------------`

	// Outputs the welcome screen
	fmt.Printf("%s\n", ascii)
}

func balanceModify(bal *float64) {
	/*
	   I.S = None or will be recieve the user balance
	   F.S = Prints the converted value to user and optionally saves the value
	*/

	// Variable declaration
	var userChoice int
	var userVal float64

	// Prints the user balance and ask for the user option
	fmt.Printf("Your Balance: $%.2f\n", *bal)
	fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\n")
	fmt.Print("Your Option: ")
	fmt.Scan(&userChoice)

	// Asks the user of the available conversion method
	for userChoice != 2 {
		if userChoice == 1 {
			fmt.Printf("\nTo add the value insert how much value will be added and to reduce the value include a minus(-) to the front of the number\n")
			fmt.Print("Your Value: ")
			fmt.Scan(&userVal)
			*bal = *bal + userVal
			fmt.Printf("\nYour new Balance: $%.2f\n", *bal)
		} else if userChoice < 1 || userChoice > 2 {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\n")
		fmt.Print("Your Option: ")
		fmt.Scan(&userChoice)
	}

	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......\n")
}

func idrToUsdConverter(bal *float64) {
	/*
	   I.S = None
	   F.S = Prints the converted value to user and optionally saves the value
	*/

	// Variable declaration
	var convertChoice int
	var userInput float64
	var saveVal string
	var rate float64 = 16.474 // The default rate value is set to the average conversion price

	// Asks the user of the available conversion method
	fmt.Printf("\nChoose Conversion Method:\n1. January - May 2025 USD/IDR Average\n2. Custom Rate\n")
	fmt.Print("Your Option: ")
	fmt.Scan(&convertChoice)

	// Error handling if the user choice is invalid
	for convertChoice != 1 && convertChoice != 2 {
		fmt.Print("Invalid Value... Please Select The Correct Options\nYour Option: ")
		fmt.Scan(&convertChoice)
	}

	// Change the rate if the user wants a custom rate for the conversion
	if convertChoice == 2 {
		fmt.Print("Set Your Prefered Rate: ")
		fmt.Scan(&rate)
		fmt.Printf("\n")
	}

	// Ask the user of the value which will be converted
	fmt.Print("How Much Would You Convert: Rp.")
	fmt.Scan(&userInput)
	fmt.Printf("\nYour Conversion Result Is: $%.2f", userInput/rate)

	// Ask the user if they want to save the converted value
	fmt.Printf("\nDo you wish to save the conversion value for the balance? y/N\n")
	fmt.Print("Your Option: ")
	fmt.Scan(&saveVal)
	if saveVal == "Yes" || saveVal == "yes" || saveVal == "y" || saveVal == "Y" {
		*bal = *bal + userInput/rate
		fmt.Printf("\nValue Saved\n")
	}

	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......\n")
}

func generateRandomizedStocks() [30]Stock {
	// Initialize base prices for stocks
	basePrices := [30]float64{
		172.34,  // AAPL
		162.12,  // ABBV
		570.12,  // ADBE
		134.22,  // AMZN
		900.45,  // AVGO
		28.45,   // BAC
		340.12,  // BRK.B
		710.78,  // COST
		162.78,  // CVX
		102.34,  // DIS
		2821.45, // GOOGL
		320.45,  // HD
		168.90,  // JNJ
		158.45,  // JPM
		61.12,   // KO
		780.34,  // LLY
		390.12,  // MA
		288.34,  // MCD
		312.67,  // META
		112.45,  // MRK
		324.56,  // MSFT
		420.45,  // NFLX
		690.12,  // NVDA
		182.34,  // PEP
		152.34,  // PG
		710.34,  // TSLA
		510.34,  // UNH
		232.12,  // V
		162.12,  // WMT
		104.56,  // XOM
	}

	// Stock ticker symbols
	stockNames := [30]string{
		"AAPL", "ABBV", "ADBE", "AMZN", "AVGO",
		"BAC", "BRK.B", "COST", "CVX", "DIS",
		"GOOGL", "HD", "JNJ", "JPM", "KO",
		"LLY", "MA", "MCD", "META", "MRK",
		"MSFT", "NFLX", "NVDA", "PEP", "PG",
		"TSLA", "UNH", "V", "WMT", "XOM",
	}

	// Create stocks array
	var stocks [30]Stock

	// Initialize each stock with randomized prices
	for i := 0; i < 30; i++ {
		stocks[i].Name = stockNames[i]
		basePrice := basePrices[i]

		// Set day 0 price as the base price
		stocks[i].Prices[0] = basePrice

		// Generate prices for days 1-5 with random fluctuations
		for day := 1; day < 6; day++ {
			// Generate random percentage change between -3% and +3%
			percentChange := (rand.Float64() * 6) - 3 // Range: -3.0 to 3.0

			// Higher volatility for certain tech stocks
			if stocks[i].Name == "TSLA" || stocks[i].Name == "NVDA" || stocks[i].Name == "META" {
				percentChange = (rand.Float64() * 10) - 5 // Range: -5.0 to 5.0
			}

			// Calculate new price based on previous day's price
			prevPrice := stocks[i].Prices[day-1]
			newPrice := prevPrice * (1 + (percentChange / 100))

			// Round to 2 decimal places
			newPrice = float64(int(newPrice*100)) / 100

			stocks[i].Prices[day] = newPrice
		}
	}

	return stocks
}

func stockProfitLossViz(bal *float64) {
	/*
	   I.S = User balance
	   F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/

	// Generate randomized stocks
	stocks := generateRandomizedStocks()

	// Variable declaration
	var userChoice, i int
	var userPortfolio Portfolio

	// Ask the user for the stock they want to invest in
	fmt.Printf("Welcome to the Stock Profit/Loss Visualization\n")
	fmt.Printf("Please choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\n")
	fmt.Print("Your Option: ")
	fmt.Scan(&userChoice)

	for userChoice != 3 {
		if userChoice == 1 {
			fmt.Printf("Available Stocks:\n")
			for i = 0; i < 30; i++ {
				// Display stocks in a more readable format with current price
				fmt.Printf("%-6s: $%.2f\n", stocks[i].Name, stocks[i].Prices[0])
			}
		} else if userChoice == 2 {
			investSellStocks(&stocks, &userPortfolio, bal)
		} else if userChoice < 1 || userChoice > 3 {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\n")
		fmt.Print("Your Option: ")
		fmt.Scan(&userChoice)
	}
	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......\n")
}

func investSellStocks(stocks *[30]Stock, userPortfolio *Portfolio, bal *float64) {
	/*
		I.S = Stocks data, User Portfolio, User Balance
		F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/

	var i, currentDay, stockOptionChoice, stockIdx, qty int
	var stockSearch string
	const MAX_PORTFOLIO_SIZE = 6

	currentDay = 1
	stockOptionChoice = 0

	// Show current balance at the start
	fmt.Printf("Current Balance: $%.2f\n", *bal)

	for stockOptionChoice != 5 && currentDay <= 5 {
		fmt.Printf("\nDay %d\n", currentDay)
		fmt.Printf("Your Portfolio:\n")

		// Display portfolio summary including total value and profit/loss
		var totalPortfolioValue float64 = 0
		var totalProfit float64 = 0

		for i = 0; i < 30; i++ {
			if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
				// Update profit and return percent for the current day
				userPortfolio.Profit[i] = (stocks[i].Prices[currentDay-1] - userPortfolio.buyPrice[i]) * float64(userPortfolio.Quantity[i])
				if userPortfolio.buyPrice[i] != 0 {
					userPortfolio.ReturnPercent[i] = (stocks[i].Prices[currentDay-1] - userPortfolio.buyPrice[i]) / userPortfolio.buyPrice[i] * 100
				} else {
					userPortfolio.ReturnPercent[i] = 0
				}

				totalPortfolioValue += stocks[i].Prices[currentDay-1] * float64(userPortfolio.Quantity[i])
				totalProfit += userPortfolio.Profit[i]

				fmt.Printf("Stock: %-6s, Buy Price: $%-8.2f, Current Price: $%-8.2f, Quantity: %-4d, Profit: $%-10.2f, Return: %6.2f%%\n",
					userPortfolio.Name[i],
					userPortfolio.buyPrice[i],
					stocks[i].Prices[currentDay-1],
					userPortfolio.Quantity[i],
					userPortfolio.Profit[i],
					userPortfolio.ReturnPercent[i])
			}
		}

		// Show portfolio summary
		if totalPortfolioValue > 0 {
			fmt.Printf("\nPortfolio Summary:\n")
			fmt.Printf("Total Portfolio Value: $%.2f\n", totalPortfolioValue)
			fmt.Printf("Total Profit/Loss: $%.2f\n", totalProfit)
			fmt.Printf("Cash Balance: $%.2f\n", *bal)
			fmt.Printf("Total Assets: $%.2f\n", totalPortfolioValue+*bal)
		}

		fmt.Print("\nSelect the following options:\n1. Previous day highest gainers\n2. Buy a stock\n3. Sell a stock\n4. Skip to next day\n5. Exit\n")
		fmt.Print("Your Option: ")
		fmt.Scan(&stockOptionChoice)

		if stockOptionChoice == 1 {
			sortHighestPrice(stocks, currentDay-1)
		} else if stockOptionChoice == 2 {
			// Check if portfolio is full
			var portfolioCount int
			for i = 0; i < 30; i++ {
				if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
					portfolioCount++
				}
			}
			if portfolioCount >= MAX_PORTFOLIO_SIZE {
				fmt.Printf("Portfolio limit reached! You cannot invest in more than %d stocks.\n", MAX_PORTFOLIO_SIZE)
				continue
			}

			// Display available stocks with current prices for easier selection
			fmt.Printf("\nAvailable Stocks (Day %d Prices):\n", currentDay)
			for i = 0; i < 5; i++ { // Show first 5 stocks
				fmt.Printf("%-6s: $%.2f\n", (*stocks)[i].Name, (*stocks)[i].Prices[currentDay-1])
			}
			fmt.Printf("... and %d more. Type stock ticker to search.\n", 30-5)

			fmt.Print("Enter stock ticker to buy: ")
			fmt.Scan(&stockSearch)
			stockIdx = searchUsingBinSearch(*stocks, stockSearch)

			if stockIdx != -1 {
				fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
				fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay-1])
				fmt.Printf("Your Balance: $%.2f\n", *bal)
				fmt.Printf("Maximum shares you can buy: %.0f\n", *bal/stocks[stockIdx].Prices[currentDay-1])

				fmt.Print("How many shares would you like to buy? ")
				fmt.Scan(&qty)

				cost := stocks[stockIdx].Prices[currentDay-1] * float64(qty)
				if qty > 0 && *bal >= cost {
					*bal -= cost
					// Check if user already owns this stock
					if userPortfolio.Name[stockIdx] == stocks[stockIdx].Name && userPortfolio.Quantity[stockIdx] > 0 {
						// Update average buy price
						totalCost := userPortfolio.buyPrice[stockIdx]*float64(userPortfolio.Quantity[stockIdx]) + cost
						userPortfolio.Quantity[stockIdx] += qty
						userPortfolio.buyPrice[stockIdx] = totalCost / float64(userPortfolio.Quantity[stockIdx])
					} else {
						userPortfolio.Name[stockIdx] = stocks[stockIdx].Name
						userPortfolio.buyPrice[stockIdx] = stocks[stockIdx].Prices[currentDay-1]
						userPortfolio.Quantity[stockIdx] = qty
					}

					// Update profit calculation
					userPortfolio.Profit[stockIdx] = (stocks[stockIdx].Prices[currentDay-1] - userPortfolio.buyPrice[stockIdx]) * float64(userPortfolio.Quantity[stockIdx])
					if userPortfolio.buyPrice[stockIdx] != 0 {
						userPortfolio.ReturnPercent[stockIdx] = (stocks[stockIdx].Prices[currentDay-1] - userPortfolio.buyPrice[stockIdx]) / userPortfolio.buyPrice[stockIdx] * 100
					} else {
						userPortfolio.ReturnPercent[stockIdx] = 0
					}

					fmt.Printf("Purchase successful! Bought %d shares of %s for $%.2f\n", qty, stocks[stockIdx].Name, cost)
					fmt.Printf("New balance: $%.2f\n", *bal)
				} else {
					fmt.Printf("Insufficient balance or invalid quantity.\n")
				}
			} else {
				fmt.Printf("Stock Not Found\n")
			}
		} else if stockOptionChoice == 3 {
			// Only show stocks the user owns
			fmt.Printf("\nYour Holdings:\n")
			var hasStocks bool = false
			for i = 0; i < 30; i++ {
				if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
					fmt.Printf("%-6s: %d shares at $%.2f (current: $%.2f)\n",
						userPortfolio.Name[i],
						userPortfolio.Quantity[i],
						userPortfolio.buyPrice[i],
						stocks[i].Prices[currentDay-1])
					hasStocks = true
				}
			}

			if !hasStocks {
				fmt.Printf("You don't own any stocks yet.\n")
				continue
			}

			fmt.Print("Enter stock ticker to sell: ")
			fmt.Scan(&stockSearch)
			stockIdx = searchUsingBinSearch(*stocks, stockSearch)

			if stockIdx != -1 && userPortfolio.Quantity[stockIdx] > 0 {
				fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
				fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay-1])
				fmt.Printf("You own %d shares, current value: $%.2f\n",
					userPortfolio.Quantity[stockIdx],
					stocks[stockIdx].Prices[currentDay-1]*float64(userPortfolio.Quantity[stockIdx]))

				fmt.Print("How many shares would you like to sell? ")
				fmt.Scan(&qty)

				if qty > 0 && qty <= userPortfolio.Quantity[stockIdx] {
					proceeds := stocks[stockIdx].Prices[currentDay-1] * float64(qty)
					profit := proceeds - (userPortfolio.buyPrice[stockIdx] * float64(qty))

					*bal += proceeds
					userPortfolio.Quantity[stockIdx] -= qty

					fmt.Printf("Sale successful! Sold %d shares of %s for $%.2f\n", qty, stocks[stockIdx].Name, proceeds)
					if profit > 0 {
						fmt.Printf("You made a profit of $%.2f on this transaction!\n", profit)
					} else if profit < 0 {
						fmt.Printf("You took a loss of $%.2f on this transaction.\n", -profit)
					}

					if userPortfolio.Quantity[stockIdx] == 0 {
						userPortfolio.Name[stockIdx] = ""
						userPortfolio.buyPrice[stockIdx] = 0
						userPortfolio.Profit[stockIdx] = 0
						userPortfolio.ReturnPercent[stockIdx] = 0
						fmt.Printf("You've sold all your shares of %s\n", stocks[stockIdx].Name)
					} else {
						fmt.Printf("You still own %d shares of %s\n", userPortfolio.Quantity[stockIdx], stocks[stockIdx].Name)
					}

					fmt.Printf("New balance: $%.2f\n", *bal)
				} else {
					fmt.Printf("Invalid quantity. You cannot sell more shares than you own.\n")
				}
			} else {
				fmt.Printf("You do not own this stock or have no shares to sell.\n")
			}
		} else if stockOptionChoice == 4 {
			if currentDay == 5 {
				fmt.Println("This is the last day of simulation.")
				fmt.Print("Press 5 to exit or any other number to continue: ")
				fmt.Scan(&stockOptionChoice)
				if stockOptionChoice != 5 {
					stockOptionChoice = 0
				}
			} else {
				fmt.Println("Skipping to the next day...")
				currentDay++
			}
		} else if stockOptionChoice != 5 {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
	}

	// Show final portfolio status if exiting
	if stockOptionChoice == 5 {
		fmt.Printf("\nFinal Portfolio Status (Day %d):\n", currentDay)
		var totalValue float64 = 0
		var initialInvestment float64 = 0

		for i = 0; i < 30; i++ {
			if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
				currentValue := stocks[i].Prices[currentDay-1] * float64(userPortfolio.Quantity[i])
				initialCost := userPortfolio.buyPrice[i] * float64(userPortfolio.Quantity[i])
				totalValue += currentValue
				initialInvestment += initialCost

				fmt.Printf("%-6s: %d shares, bought at $%.2f, final price $%.2f, P/L: $%.2f (%.2f%%)\n",
					userPortfolio.Name[i],
					userPortfolio.Quantity[i],
					userPortfolio.buyPrice[i],
					stocks[i].Prices[currentDay-1],
					currentValue-initialCost,
					(stocks[i].Prices[currentDay-1]/userPortfolio.buyPrice[i]-1)*100)
			}
		}

		if totalValue > 0 {
			fmt.Printf("\nTotal Portfolio Value: $%.2f\n", totalValue)
			fmt.Printf("Total Initial Investment: $%.2f\n", initialInvestment)
			fmt.Printf("Total Profit/Loss: $%.2f (%.2f%%)\n",
				totalValue-initialInvestment,
				(totalValue/initialInvestment-1)*100)
			fmt.Printf("Cash Balance: $%.2f\n", *bal)
			fmt.Printf("Total Assets: $%.2f\n", totalValue+*bal)
		} else {
			fmt.Printf("\nYou have no stocks in your portfolio.\n")
			fmt.Printf("Cash Balance: $%.2f\n", *bal)
		}
	}
}

func sortHighestPrice(stocks *[30]Stock, dayIndex int) {
	/*
	   I.S = Stocks data
	   F.S = Sorts the stocks data by the highest price
	*/

	// Create a copy of the stocks array
	var sortedStocks [30]Stock
	for i := 0; i < 30; i++ {
		sortedStocks[i] = (*stocks)[i]
	}

	// Calculate price changes for analysis
	type StockChange struct {
		Name          string
		Price         float64
		PriceChange   float64
		PercentChange float64
	}

	var changes [30]StockChange
	for i := 0; i < 30; i++ {
		changes[i].Name = sortedStocks[i].Name
		changes[i].Price = sortedStocks[i].Prices[dayIndex]

		// Calculate change if not the first day
		if dayIndex > 0 {
			changes[i].PriceChange = sortedStocks[i].Prices[dayIndex] - sortedStocks[i].Prices[dayIndex-1]
			changes[i].PercentChange = (changes[i].PriceChange / sortedStocks[i].Prices[dayIndex-1]) * 100
		}
	}

	// Sort the copy by the highest price using selection sort
	var i, min, pass int
	var tempStock Stock
	for pass = 0; pass < 30; pass++ {
		min = pass
		for i = pass + 1; i < 30; i++ {
			if sortedStocks[min].Prices[dayIndex] < sortedStocks[i].Prices[dayIndex] {
				min = i
			}
		}
		tempStock = sortedStocks[pass]
		sortedStocks[pass] = sortedStocks[min]
		sortedStocks[min] = tempStock
	}

	// Display the top 5 highest price stocks with more details
	fmt.Printf("Top 5 highest priced stocks on day %d:\n", dayIndex+1)
	fmt.Printf("%-6s | %-10s | %-10s | %-10s\n", "Symbol", "Price", "Change", "% Change")
	fmt.Printf("--------------------------------------------\n")
	for i = 0; i < 5; i++ {
		// Find the stock change info
		var changeIdx int
		for j := 0; j < 30; j++ {
			if changes[j].Name == sortedStocks[i].Name {
				changeIdx = j
				break
			}
		}

		// Display with price change if not the first day
		if dayIndex > 0 {
			fmt.Printf("%-6s | $%-9.2f | $%-9.2f | %-10.2f%%\n",
				sortedStocks[i].Name,
				sortedStocks[i].Prices[dayIndex],
				changes[changeIdx].PriceChange,
				changes[changeIdx].PercentChange)
		} else {
			fmt.Printf("%-6s | $%-9.2f | %-10s | %-10s\n",
				sortedStocks[i].Name,
				sortedStocks[i].Prices[dayIndex],
				"N/A", "N/A")
		}
	}
}

func searchUsingBinSearch(stocks [30]Stock, stockName string) int {
	/*
	   I.S = Stocks data
	   F.S = Returns the index of the stock if found or -1 if not found
	*/

	// Variable declaration
	var low, high, mid int

	// Binary search for the stock
	low = 0
	high = 29
	for low <= high {
		mid = (low + high) / 2
		if stocks[mid].Name == stockName {
			return mid
		} else if stocks[mid].Name < stockName {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
