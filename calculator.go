package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"project/calculator/evaluator" // Import evaluator package
)

func main() {
	// Initialize the application
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	// Input and output fields
	expression := widget.NewEntry()
	expression.SetPlaceHolder("Enter expression")

	resultLabel := widget.NewLabel("Result: ")

	// Add buttons for digits and operations
	buttons := []string{
		"7", "8", "9", "/", "C",
		"4", "5", "6", "*", "(",
		"1", "2", "3", "-", ")",
		"0", ".", "=", "+", "%",
	}

	grid := container.NewGridWithColumns(5)
	for _, b := range buttons {
		btn := b // Capture the loop variable
		grid.Add(widget.NewButton(btn, func() {
			if btn == "=" {
				res, err := evaluator.EvaluateExpression(expression.Text) // Call from evaluator package
				if err != nil {
					resultLabel.SetText("Error: " + err.Error())
				} else {
					resultLabel.SetText(fmt.Sprintf("Result: %v", res))
				}
			} else if btn == "C" {
				expression.SetText("")
				resultLabel.SetText("Result: ")
			} else {
				expression.SetText(expression.Text + btn)
			}
		}))
	}

	// Layout for the app
	content := container.NewVBox(
		expression,
		resultLabel,
		grid,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}