package main

import "fmt"

type Stock struct {
	Name   string
	Prices [6]float64
}

type Portfolio struct {
	Name          [6]string
	buyPrice      [6]float64
	Quantity      [6]int
	Profit        [6]float64
	ReturnPercent [6]float64
}

func main() {
	var choice int
	var balance float64
	var selectionMethod, searchMethod int = 1, 1

	welcomeScreen()
	fmt.Scan(&choice)

	for choice != 5 {
		if choice == 1 {
			idrToUsdConverter(&balance)
		} else if choice == 2 {
			balanceModify(&balance)
		} else if choice == 3 {
			stockProfitLossViz(&balance, selectionMethod, searchMethod)
		} else if choice == 4 {
			options(&selectionMethod, &searchMethod)
		} else if choice < 0 || choice > 5 {
			fmt.Printf("Invalid Choice!\n")
		}
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
	var rate float64 = 16.474 // The default rate value is set to the average conversion price

	fmt.Printf("\nChoose Conversion Method:\n1. January - May 2025 USD/IDR Average\n2. Custom Rate\nYour Option: ")
	fmt.Scan(&convertChoice)

	for convertChoice != 1 && convertChoice != 2 {
		fmt.Printf("Invalid Value... Please Select The Correct Options\nYour Option: ")
		fmt.Scan(&convertChoice)
	}

	if convertChoice == 2 {
		fmt.Printf("Set Your Prefered Rate: ")
		fmt.Scan(&rate)
		fmt.Printf("\n")
	}

	fmt.Printf("How Much Would You Convert: Rp.")
	fmt.Scan(&userInput)
	fmt.Printf("\nYour Conversion Result Is: $%.2f", userInput/rate)

	fmt.Printf("\nDo you wish to save the conversion value for the balance? y/N\nYour Option: ")
	fmt.Scan(&saveVal)
	if saveVal == "Yes" || saveVal == "yes" || saveVal == "y" || saveVal == "Y" {
		*bal = *bal + userInput/rate
		fmt.Printf("\nValue Saved")
	}
	fmt.Printf("Returning to the main screen.......")
}

func balanceModify(bal *float64) {
	/*
	   I.S = None or will be recieve the user balance
	   F.S = Prints the converted value to user and optionally saves the value
	*/

	var userChoice int
	var userVal float64

	fmt.Printf("Your Balance: $%.2f\n", *bal)
	fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: \n")
	fmt.Scan(&userChoice)

	for userChoice != 2 {
		if userChoice == 1 {
			fmt.Printf("\nEnter a value to add or subtract (use '-' for reduction):\nYour Option: ")
			fmt.Scan(&userVal)
			*bal = *bal + userVal
			fmt.Printf("\nYour new Balance: $%.2f\n", *bal)

		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: ")
		fmt.Scan(&userChoice)
	}
	fmt.Printf("Returning to the main screen.......")
}

func stockProfitLossViz(bal *float64, selectionMethod, searchMethod int) {
	/*
	   I.S = User balance
	   F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/
	var stocks [30]Stock
	stocks = [30]Stock{
		{"AAPL", [6]float64{172.34, 168.21, 175.89, 170.56, 173.12, 174.50}},
		{"ABBV", [6]float64{162.12, 164.45, 163.78, 165.90, 162.67, 166.10}},
		{"ADBE", [6]float64{570.12, 572.45, 571.78, 573.90, 570.67, 574.20}},
		{"AMZN", [6]float64{134.22, 137.45, 136.78, 138.90, 135.67, 139.20}},
		{"AVGO", [6]float64{900.45, 902.12, 901.67, 903.34, 900.90, 904.00}},
		{"BAC", [6]float64{28.45, 29.12, 28.67, 29.34, 28.90, 29.50}},
		{"BRK.B", [6]float64{340.12, 342.45, 341.78, 343.90, 340.67, 344.10}},
		{"COST", [6]float64{710.78, 712.12, 711.67, 713.45, 710.89, 714.20}},
		{"CVX", [6]float64{162.78, 164.12, 163.67, 165.45, 162.89, 165.90}},
		{"DIS", [6]float64{102.34, 104.12, 103.67, 105.45, 102.89, 105.90}},
		{"GOOGL", [6]float64{2821.45, 2805.67, 2830.12, 2810.34, 2825.78, 2832.00}},
		{"HD", [6]float64{320.45, 322.12, 321.67, 323.34, 320.90, 324.00}},
		{"JNJ", [6]float64{168.90, 170.45, 169.78, 171.12, 170.34, 171.80}},
		{"JPM", [6]float64{158.45, 160.12, 159.67, 161.34, 158.90, 162.00}},
		{"KO", [6]float64{61.12, 62.45, 61.78, 63.90, 61.67, 64.10}},
		{"LLY", [6]float64{780.34, 782.12, 781.67, 783.45, 780.89, 784.00}},
		{"MA", [6]float64{390.12, 392.45, 391.78, 393.90, 390.67, 394.20}},
		{"MCD", [6]float64{288.34, 290.12, 289.67, 291.45, 288.89, 292.00}},
		{"META", [6]float64{312.67, 315.12, 310.45, 317.89, 313.56, 318.20}},
		{"MRK", [6]float64{112.45, 114.12, 113.67, 115.34, 112.90, 115.80}},
		{"MSFT", [6]float64{324.56, 329.12, 327.45, 330.78, 328.90, 331.10}},
		{"NFLX", [6]float64{420.45, 422.12, 421.67, 423.34, 420.90, 424.00}},
		{"NVDA", [6]float64{690.12, 700.45, 695.78, 702.34, 698.56, 703.90}},
		{"PEP", [6]float64{182.34, 184.12, 183.67, 185.45, 182.89, 185.90}},
		{"PG", [6]float64{152.34, 154.12, 153.67, 155.45, 152.89, 155.90}},
		{"TSLA", [6]float64{710.34, 705.12, 715.67, 712.45, 713.89, 716.00}},
		{"UNH", [6]float64{510.34, 512.12, 511.67, 513.45, 510.89, 514.00}},
		{"V", [6]float64{232.12, 234.45, 233.78, 235.90, 232.67, 236.10}},
		{"WMT", [6]float64{162.12, 164.45, 163.78, 165.90, 162.67, 166.10}},
		{"XOM", [6]float64{104.56, 106.12, 105.78, 107.45, 104.90, 107.80}},
	}

	var userChoice, i int
	var userPortfolio Portfolio

	for userChoice != 3 {
		if userChoice == 1 {
			fmt.Printf("Available Stocks:\n")
			for i = 0; i < 30; i += 5 {
				fmt.Printf("%s %s %s %s %s\n", stocks[i].Name, stocks[i+1].Name, stocks[i+2].Name, stocks[i+3].Name, stocks[i+4].Name)
			}
		} else if userChoice == 2 {
			investSellStocks(&stocks, &userPortfolio, bal, selectionMethod, searchMethod)
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\nYour Option: ")
		fmt.Scan(&userChoice)
	}
	fmt.Printf("Returning to the main screen.......")
}

func investSellStocks(stocks *[30]Stock, userPortfolio *Portfolio, bal *float64, selectionMethod, searchMethod int) {
	/*
		I.S = Stocks data, User Portfolio, User Balance
		F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/

	var i, currentDay, stockOptionChoice, stockIdx, qty, portfolioCount int
	var stockSearch string
	var cost, proceeds float64

	currentDay = 1
	for currentDay <= 5 {
		fmt.Printf("\nTrading Day %d\n", currentDay)
		fmt.Printf("Your Portfolio:\n")
		for i = 0; i < 6; i++ {
			if userPortfolio.Quantity[i] > 0 {
				if currentDay > 1 { // Only calculate profit from day 2 onwards
					userPortfolio.Profit[i] = (stocks[i].Prices[currentDay] - userPortfolio.buyPrice[i]) * float64(userPortfolio.Quantity[i])
					userPortfolio.ReturnPercent[i] = (userPortfolio.Profit[i] / userPortfolio.buyPrice[i]) * 100
					fmt.Printf("Stock: %s, Buy Price: $%.2f, Quantity: %d, Profit: $%.2f, Return Percent: %.2f%%\n",
						userPortfolio.Name[i], userPortfolio.buyPrice[i], userPortfolio.Quantity[i], userPortfolio.Profit[i], userPortfolio.ReturnPercent[i])
				} else {
					fmt.Printf("Stock: %s, Buy Price: $%.2f, Quantity: %d\n",
						userPortfolio.Name[i], userPortfolio.buyPrice[i], userPortfolio.Quantity[i]) // On day 1, only show the stock holdings without profit
				}
			}
		}

		// Trading menu loop for current day
		stockOptionChoice = 0
		for stockOptionChoice != 4 && stockOptionChoice != 5 {
			fmt.Print("\nSelect the following options:\n1. Previous day highest gainers\n2. Buy a stock\n3. Sell a stock\n4. Skip to next day\n5. Exit\nYour Option: ")
			fmt.Scan(&stockOptionChoice)
			if stockOptionChoice == 1 {
				if selectionMethod == 1 {
					sortHighestPriceSelection(*stocks, currentDay-1)
				} else {
					sortHighestPriceInsertion(*stocks, currentDay-1)
				}
			} else if stockOptionChoice == 2 {
				if userPortfolio.Quantity[stockIdx] == 0 && portfolioCount >= 6 {
					fmt.Println("Portfolio limit reached! You cannot invest in more than 6 different stocks.")
					return
				}
				fmt.Printf("Search a stock to buy:\n")
				fmt.Scan(&stockSearch)

				if searchMethod == 1 {
					stockIdx = searchUsingSeqSearch(*stocks, stockSearch)
				} else {
					stockIdx = searchUsingBinSearch(*stocks, stockSearch)
				}
				if stockIdx != -1 {
					fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
					fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay])
					fmt.Printf("How many shares would you like to buy?\nYour Option: ")
					fmt.Scan(&qty)
					cost = stocks[stockIdx].Prices[currentDay] * float64(qty)
					if qty > 0 && *bal >= cost {
						*bal -= cost
						if userPortfolio.Quantity[stockIdx] == 0 {
							userPortfolio.Name[stockIdx] = stocks[stockIdx].Name
							portfolioCount++
						}
						userPortfolio.buyPrice[stockIdx] = stocks[stockIdx].Prices[currentDay]
						userPortfolio.Quantity[stockIdx] += qty
						fmt.Printf("Purchase successful! New balance: $%.2f\n", *bal)

					} else {
						fmt.Printf("Insufficient balance or invalid quantity.\n")
					}
				} else {
					fmt.Printf("Stock Not Found\n")
				}
			} else if stockOptionChoice == 3 {
				fmt.Printf("Search a stock to sell:\n")
				fmt.Scan(&stockSearch)
				stockIdx = searchUsingSeqSearch(*stocks, stockSearch)
				if stockIdx != -1 && userPortfolio.Quantity[stockIdx] > 0 {
					fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
					fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay])
					fmt.Printf("How many shares would you like to sell? (You own %d)\nYour Option: ", userPortfolio.Quantity[stockIdx])
					fmt.Scan(&qty)
					if qty > 0 && qty <= userPortfolio.Quantity[stockIdx] {
						proceeds = stocks[stockIdx].Prices[currentDay] * float64(qty)
						*bal += proceeds
						userPortfolio.Quantity[stockIdx] -= qty
						if userPortfolio.Quantity[stockIdx] == 0 {
							userPortfolio.Name[stockIdx] = ""
							userPortfolio.buyPrice[stockIdx] = 0
							userPortfolio.Profit[stockIdx] = 0
							userPortfolio.ReturnPercent[stockIdx] = 0
						}
						fmt.Printf("Sale successful! New balance: $%.2f\n", *bal)
					} else {
						fmt.Printf("Invalid quantity.\n")
					}
				} else {
					fmt.Printf("You do not own this stock or have no shares to sell.\n")
				}
			} else if stockOptionChoice == 4 {
				fmt.Println("Skipping to the next day...")
				currentDay++
			} else if stockOptionChoice != 5 {
				fmt.Printf("Invalid Value... Please Select The Correct Options\n")
			}
		}
	}
}

func sortHighestPriceSelection(stocks [30]Stock, n int) {
	/*
	   I.S = Stocks data and current trading day
	   F.S = Sorts and displays the stocks data by the highest price from the previous day
	   Note: When trading starts (day 1), it shows day 0 prices
	        For subsequent days, it shows the previous trading day's prices
	*/

	// Sort the copy by the highest price of the previous day
	var i, min, pass int
	var temp Stock

	for pass = 0; pass < 30; pass++ {
		min = pass
		for i = pass + 1; i < 30; i++ {
			if stocks[min].Prices[n] < stocks[i].Prices[n] {
				min = i
			}
		}
		temp = stocks[pass]
		stocks[pass] = stocks[min]
		stocks[min] = temp
	}

	// Display the top 5 highest price stocks from the previous day
	fmt.Printf("Top 5 highest prices from previous day (Day %d):\n", n)
	for i = 0; i < 5; i++ {
		fmt.Printf("%s: $%.2f\n", stocks[i].Name, stocks[i].Prices[n])
	}
}

func sortHighestPriceInsertion(stocks [30]Stock, n int) {
	/*
	   I.S = Stocks data and current trading day
	   F.S = Sorts and displays the stocks data by the highest price from the previous day
	   Note: When trading starts (day 1), it shows day 0 prices
	        For subsequent days, it shows the previous trading day's prices
	*/

	// Sort the copy by the highest price of the previous day
	var pass, j, i int
	var temp Stock

	for pass = 1; pass < 30; pass++ {
		j = pass
		temp = stocks[pass]
		for j > 0 && i >= 0 && stocks[i].Prices[n] < temp.Prices[n] {
			stocks[j] = stocks[i]
			j--
		}
		stocks[j] = temp
	}

	// Display the top 5 highest price stocks from the previous day
	fmt.Printf("Top 5 highest prices from previous day (Day %d):\n", n)
	for i = 0; i < 5; i++ {
		fmt.Printf("%s: $%.2f\n", stocks[i].Name, stocks[i].Prices[n])
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

func searchUsingSeqSearch(stocks [30]Stock, stockName string) int {
	/*
	   I.S = Stocks data
	   F.S = Returns the index of the stock if found or -1 if not found
	*/

	var i int

	for i = 0; i <= 29; i++ {
		if stocks[i].Name == stockName {
			return i
		}
	}
	return -1
}

func options(sortMethod, searchMethod *int) {
	/*
		I.S = selectionMethod, searchMethod with the default selection sort and sequential search
		F.S = user set their prefered method
	*/

	var untilExit bool = true
	var userChoice int

	for untilExit {
		fmt.Printf("Options Menu:\n1. Sorting Method\n2. Search Method\n3. Exit\nYour Option: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Printf("Select the sorting method(default:Selection Sort):\n1. Selection Sort\n2. Insertion Sort\nYour Option: ")
			fmt.Scan(sortMethod)
		} else if userChoice == 2 {
			fmt.Printf("Select the search method(default:Sequential Search):\n1. Sequential Search\n2. Binary Search\nYour Option: ")
			fmt.Scan(searchMethod)
		} else if userChoice == 3 {
			untilExit = false
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
	}
	fmt.Printf("Returning to the main screen.......")
}
