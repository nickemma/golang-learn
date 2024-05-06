package main

import "fmt"

//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Access Granted")
}

func accessDenied() {
	fmt.Println("Access Denied")
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Guest

	if role == Admin || role == Manager {
		// Access granted for Admin and Manager at any time.
		accessGranted()
	} else if role == Contractor {
		// Access granted for Contractor on weekends.
		if today == Saturday || today == Sunday {
			accessGranted()
		} else {
			accessDenied()
		}
	} else if role == Member {
		// Access granted for Member on weekdays.
		if today >= Monday && today <= Friday {
			accessGranted()
		} else {
			accessDenied()
		}
	} else if role == Guest {
		// Access granted for Guest on Mondays, Wednesdays, and Fridays.
		if today == Monday || today == Wednesday || today == Friday {
			accessGranted()
		} else {
			accessDenied()
		}
	} else {
		// Access denied for other roles.
		accessDenied()
	}
}
