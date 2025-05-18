package main

import (
	"fmt"
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

	// Shows the title screen and asks the user choice
	welcomeScreen()
	fmt.Scan(&choice)

	// Redirect the user option to the corresponding function and repeats the title screen until the user option is 3 (exit)
	for choice != 4 {
		if choice == 1 {
			idrToUsdConverter(&balance)
		} else if choice == 2 {
			balanceModify(&balance)
		} else if choice == 3 {
			stockProfitLossViz(&balance)
		} else if choice < 0 || choice > 3 {
			fmt.Printf("Invalid Choice!\n")
		}
		welcomeScreen()
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
Welcome to Stock Tracker (US ONLY)
============================================================
1. Indonesian Rupiah To US Dollar Conversion
2. Check / Increase / Decrease the balance in USD
3. Stock Vizualization
4. Exit
------------------------------------------------------------
Your Option: `

	// Outputs the welcome screen
	fmt.Printf("%s", ascii)
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
	fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: \n")
	fmt.Scan(&userChoice)

	// Asks the user of the available conversion method
	for userChoice != 2 {
		if userChoice == 1 {
			fmt.Printf("\nTo add the value insert how much value will be added and to reduce the value include a minus(-) to the front of the number\nYour Option: ")
			fmt.Scan(&userVal)
			*bal = *bal + userVal
			fmt.Printf("\nYour new Balance: $%.2f\n", *bal)

		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. Increase or Decrease balance\n2. Exit\nYour Option: ")
		fmt.Scan(&userChoice)
	}

	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......")
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
	fmt.Printf("\nChoose Conversion Method:\n1. January - May 2025 USD/IDR Average\n2. Custom Rate\nYour Option: ")
	fmt.Scan(&convertChoice)

	// Error handling if the user choice is invalid
	for convertChoice != 1 && convertChoice != 2 {
		fmt.Printf("Invalid Value... Please Select The Correct Options\nYour Option: ")
		fmt.Scan(&convertChoice)
	}

	// Change the rate if the user wants a custom rate for the conversion
	if convertChoice == 2 {
		fmt.Printf("Set Your Prefered Rate: ")
		fmt.Scan(&rate)
		fmt.Printf("\n")
	}

	// Ask the user of the value which will be converted
	fmt.Printf("How Much Would You Convert: Rp.")
	fmt.Scan(&userInput)
	fmt.Printf("\nYour Conversion Result Is: $%.2f", userInput/rate)

	// Ask the user if they want to save the converted value
	fmt.Printf("\nDo you wish to save the conversion value for the balance? y/N\nYour Option: ")
	fmt.Scan(&saveVal)
	if saveVal == "Yes" || saveVal == "yes" || saveVal == "y" || saveVal == "Y" {
		*bal = *bal + userInput/rate
		fmt.Printf("\nValue Saved")
	}

	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......")
}

func stockProfitLossViz(bal *float64) {
	/*
	   I.S = User balance
	   F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/

	// Stocks data alphabetically sorted
	stocks := [30]Stock{
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

	// Variable declaration
	var userChoice, i int
	var userPortfolio Portfolio

	// Ask the user for the stock they want to invest in
	fmt.Printf("Welcome to the Stock Profit/Loss Visualization\n")
	fmt.Printf("Please choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\nYour Option: ")
	fmt.Scan(&userChoice)

	for userChoice != 3 {
		if userChoice == 1 {
			fmt.Printf("Available Stocks:\n")
			for i = 0; i < 30; i++ {
				fmt.Println(stocks[i].Name)
			}
		} else if userChoice == 2 {
			investSellStocks(&stocks, &userPortfolio, bal)
		} else {
			fmt.Printf("Invalid Value... Please Select The Correct Options\n")
		}
		fmt.Printf("Please choose the following options:\n1. View Stocks\n2. Invest/Sell Stocks\n3. Exit\nYour Option: ")
		fmt.Scan(&userChoice)
	}
	// Text to show the program will be back to the title screen
	fmt.Printf("Returning to the main screen.......")
}

func investSellStocks(stocks *[30]Stock, userPortfolio *Portfolio, bal *float64) {
	/*
		I.S = Stocks data, User Portfolio, User Balance
		F.S = Shows the user the list of stocks and the profit/loss if the user decides to invest
	*/

	var i, currentDay, stockOptionChoice, stockIdx, qty int
	var stockSearch string

	currentDay = 1
	stockOptionChoice = 0
	for stockOptionChoice != 5 && currentDay < 6 {
		fmt.Printf("\nDay %d\n", currentDay)
		fmt.Printf("Your Portfolio:\n")
		for i = 0; i < 30; i++ {
			if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
				// Update profit and return percent for the current day
				userPortfolio.Profit[i] = (stocks[i].Prices[currentDay-1] - userPortfolio.buyPrice[i]) * float64(userPortfolio.Quantity[i])
				if userPortfolio.buyPrice[i] != 0 {
					userPortfolio.ReturnPercent[i] = (userPortfolio.Profit[i] / userPortfolio.buyPrice[i]) * 100
				} else {
					userPortfolio.ReturnPercent[i] = 0
				}
				fmt.Printf("Stock: %s, Buy Price: $%.2f, Quantity: %d, Profit: $%.2f, Return Percent: %.2f%%\n", userPortfolio.Name[i], userPortfolio.buyPrice[i], userPortfolio.Quantity[i], userPortfolio.Profit[i], userPortfolio.ReturnPercent[i])
			}
		}
		fmt.Print("Select the following options:\n1. Previous day highest gainers\n2. Buy a stock\n3. Sell a stock\n4. Skip to next day\n5. Exit\nYour Option: ")
		fmt.Scan(&stockOptionChoice)
		if stockOptionChoice == 1 {
			sortHighestPrice(stocks, currentDay-1)
		} else if stockOptionChoice == 2 {
			// Check if portfolio is full
			var portfolioCount int
			for i = 0; i < 6; i++ {
				if userPortfolio.Name[i] != "" && userPortfolio.Quantity[i] > 0 {
					portfolioCount++
				}
			}
			if portfolioCount >= 6 {
				fmt.Println("Portfolio limit reached! You cannot invest in more than 6 stocks.")
				return
			}
			fmt.Printf("Search a stock to buy:\n")
			fmt.Scan(&stockSearch)
			stockIdx = searchUsingBinSearch(*stocks, stockSearch)
			if stockIdx != -1 {
				fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
				fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay-1])
				fmt.Printf("How many shares would you like to buy?\nYour Option: ")
				fmt.Scan(&qty)
				cost := stocks[stockIdx].Prices[currentDay-1] * float64(qty)
				if qty > 0 && *bal >= cost {
					*bal -= cost
					userPortfolio.Name[stockIdx] = stocks[stockIdx].Name
					userPortfolio.buyPrice[stockIdx] = stocks[stockIdx].Prices[currentDay-1]
					userPortfolio.Quantity[stockIdx] += qty
					if currentDay > 0 {
						userPortfolio.Profit[stockIdx] = (stocks[stockIdx].Prices[currentDay-1] - userPortfolio.buyPrice[stockIdx]) * float64(userPortfolio.Quantity[stockIdx])
						userPortfolio.ReturnPercent[stockIdx] = (userPortfolio.Profit[stockIdx] / userPortfolio.buyPrice[stockIdx]) * 100
					} else {
						userPortfolio.Profit[stockIdx] = 0
						userPortfolio.ReturnPercent[stockIdx] = 0
					}
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
			stockIdx = searchUsingBinSearch(*stocks, stockSearch)
			if stockIdx != -1 && userPortfolio.Quantity[stockIdx] > 0 {
				fmt.Printf("Stock Found: %s\n", stocks[stockIdx].Name)
				fmt.Printf("Current Price: $%.2f\n", stocks[stockIdx].Prices[currentDay-1])
				fmt.Printf("How many shares would you like to sell? (You own %d)\nYour Option: ", userPortfolio.Quantity[stockIdx])
				fmt.Scan(&qty)
				if qty > 0 && qty <= userPortfolio.Quantity[stockIdx] {
					proceeds := stocks[stockIdx].Prices[currentDay-1] * float64(qty)
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

func sortHighestPrice(stocks *[30]Stock, n int) {
	/*
	   I.S = Stocks data
	   F.S = Sorts the stocks data by the highest price
	*/

	// Create a copy of the stocks array
	var sortedStocks [30]Stock
	for i := 0; i < 30; i++ {
		sortedStocks[i] = (*stocks)[i]
	}
	// Sort the copy by the highest price
	var i, min, pass int
	var temp Stock
	for pass = 0; pass < 30; pass++ {
		min = pass
		for i = pass + 1; i < 30; i++ {
			if sortedStocks[min].Prices[n] < sortedStocks[i].Prices[n] {
				min = i
			}
		}
		temp = sortedStocks[pass]
		sortedStocks[pass] = sortedStocks[min]
		sortedStocks[min] = temp
	}
	// Display the top 5 highest price stocks
	fmt.Printf("The highest price on the previous day are:\n")
	for i = 0; i < 5; i++ {
		fmt.Printf("%s: $%.2f\n", sortedStocks[i].Name, sortedStocks[i].Prices[n])
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
