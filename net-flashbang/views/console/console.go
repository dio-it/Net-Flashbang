package console

import (
	"fmt"
)

// PrintMenu prints the menu to the console
func PrintMenu() {
	fmt.Println(`
	###########################################
	#******* WELCOME TO OUR LIBRARY ***********
	#******* CHOOSE YOUR OPTION BELOW *********
	# 1. ADD ANOTHER BOOK IN LIBRARY (TO BE IMPLEMENTED)
	# 2. REMOVE A BOOK FROM A LIBRARY (TO BE IMPLEMENTED)
	# 3. CHECK AVAILABILITY (TO BE IMPLEMENTED)
	# 4. LEND A BOOK (TO BE IMPLEMENTED)
	# 5. RETURN A BOOK (TO BE IMPLEMENTED)
	# 6. VIEW ALL BOOKS
	#
	# c. CLEAR VIEW AND PRINT MENU
	# q. TERMINATE BOOK LIBRARY APP
	`)
}

func Test() {
	fmt.Println("Hallo user")
}
