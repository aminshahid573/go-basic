package main

import (
	"fmt"

	"github.com/aminshahid573/go-basic/auth"
	"github.com/aminshahid573/go-basic/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("aminShahid", "secret")

	session := auth.GetSession()
	fmt.Println("Session", session)

	user := user.User{
		Email: "User@gmail.com",
		Name:  "Shahid Amin",
	}

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground = red.Add(color.Bold)
	whiteBackground.Println("User is Looged in as", user.Name)

	// Mix with RGB color

}
