package main

import "time"
var (
	exit = false
	result string
)

func main() {
	for !exit {
		// Clear Result
		result = ""
		//clearDisplay()
		result = time.Now().Local().Format("01-02       1504") + "\n"
		

		// Result is max 16x16 characters
		print(result)
		clearDisplay()
		time.Sleep(1*time.Second)
	}
}

func clearDisplay() {
	// Modify this for the exact amount of \n needed
	//				1 2 3 4 5 6 7 8 9 0 2 3 4 5 6 7 
	cleanScreen := "\n\n\n\n\n\n\n\n\n\n\n\n\n\n"
	print(cleanScreen)
}