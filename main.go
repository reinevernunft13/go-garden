package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/widget"
)

// creates struct with config settings
type Config struct {
	App        fyne.App // defines what GUI fyne will employ
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	MainWindow fyne.Window // stores reference to main window

}

var myApp Config // var to store the app's config
func main() {
	//Creates fyne app
	fyneApp := app.NewWithID("cat.cibernarium.go-garden")
	myApp.App = fyneApp
	//Invokes logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)        // creates data logger
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile) // creates error logger
	//DB connection

	//Creates DB repo

	//Creates and defines fyne's window
	myApp.MainWindow = fyneApp.NewWindow("go-run app")
	myApp.MainWindow.Resize(fyne.NewSize(800, 500)) // sets window size
	myApp.MainWindow.SetFixedSize(true)             // has a fixed size
	myApp.MainWindow.SetMaster()

	myApp.makeUI() // invokes an external fn where the UI will be designed
	//Shows and runs the app
	myApp.MainWindow.ShowAndRun()
}
