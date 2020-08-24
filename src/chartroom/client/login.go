package main

import "fmt"

func login(userID int, userPWD string) (err error) {
	fmt.Printf("userId=%d userpwd =%s \n", userID, userPWD)

	return nil
}
