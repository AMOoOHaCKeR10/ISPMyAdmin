package main

import (
	"ispmyadmin/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// Removed OpenGL bindings import as it caused build issues
)

func mainGUI() {
	// Initialize the app
	myApp := app.New()
	myWindow := myApp.NewWindow("ISPmyAdmin")

	// Title
	title := widget.NewLabel("Welcome to ISPmyAdmin")

	// Add User Form
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")
	addUserButton := widget.NewButton("Add User", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		if username == "" || password == "" {
			widget.ShowPopUp(widget.NewLabel("Username and password cannot be empty"), myWindow.Canvas())
			return
		}
		err := server.AddUser(username, password)
		if err != nil {
			widget.ShowPopUp(widget.NewLabel("Failed to add user: "+err.Error()), myWindow.Canvas())
		} else {
			widget.ShowPopUp(widget.NewLabel("User added successfully!"), myWindow.Canvas())
		}
	})

	// View Users Button
	viewUsersButton := widget.NewButton("View Users", func() {
		users := server.GetAllUsers()
		userList := ""
		for _, user := range users {
			userList += user.Username + "\n"
		}
		widget.ShowPopUp(widget.NewLabel("Users:\n"+userList), myWindow.Canvas())
	})

	// Layout
	content := container.NewVBox(
		title,
		widget.NewLabel("Add a New User:"),
		usernameEntry,
		passwordEntry,
		addUserButton,
		viewUsersButton,
	)

	// Set content and show the window
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
