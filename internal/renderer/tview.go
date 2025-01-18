package renderer

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TviewRenderer(cp Parser) {
	app := tview.NewApplication().EnablePaste(true)

	left := tview.NewFlex().SetDirection(tview.FlexRow)
	left.SetBackgroundColor(tcell.ColorDefault)

	right := tview.NewFlex()
	right.SetBackgroundColor(tcell.ColorDefault)

	separator := tview.NewTextView().SetText(strings.Repeat("│\n", 8))
	separator.SetBackgroundColor(tcell.ColorDefault)
	separator.SetTextColor(tcell.ColorWhite)

	inputField := tview.NewInputField().
		SetPlaceholder("* * * * *").
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(tcell.ColorWhite).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorDefault)).
		SetLabelStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	inputField.SetBackgroundColor(tcell.ColorDefault)
	inputField.SetAcceptanceFunc(func(text string, _ rune) bool {
		inpLen := len(strings.Split(strings.Trim(text, " "), " "))
		return text != " " && inpLen <= 5 && !strings.HasSuffix(text, "  ")
	})

	hreadableStr := tview.NewTextView().SetTextColor(tcell.ColorYellow.TrueColor())
	hreadableStr.SetBackgroundColor(tcell.ColorDefault)

	hintsView := tview.NewTextView().SetTextColor(tcell.ColorOrangeRed.TrueColor())
	hintsView.SetBackgroundColor(tcell.ColorDefault)

	left.AddItem(inputField, 0, 4, true)
	left.AddItem(hintsView, 0, 96, false)

	right.AddItem(hreadableStr, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(left, 0, 1, true).
		AddItem(separator, 1, 0, false).
		AddItem(right, 0, 4, false)
	flex.SetBackgroundColor(tcell.ColorDefault)

	updateExpr := func(text string) {
		hreadableStr.Clear()
		hintsView.Clear()

		text = strings.Trim(text, " ")
		if text == "" {
			return
		}

		splitStr := strings.Split(text, " ")
		inpLen := len(splitStr)
		if inpLen > 5 {
			hreadableStr.SetText("INVALID CRON EXPRESSION")
			return
		}

		currElem := splitStr[inpLen-1]
		padding := len(text) - len(currElem)

		hintsView.SetText(cp.GetHints(padding, inpLen-1))

		err := cp.SetExpr(text)
		if err != nil {
			hreadableStr.SetText("INVALID CRON EXPRESSION")
			return
		}
		desc, err := cp.HumanReadableStr()
		if err != nil {
			hreadableStr.SetText("INVALID CRON EXPRESSION")
			return
		}
		newStr := strings.Builder{}
		newStr.Grow(1024)
		newStr.WriteString(desc + "\n\n")
		newStr.WriteString("Next 3 Iterations:\n\n")
		iterations, err := cp.NextInstances(3)
		if err != nil {
			hreadableStr.SetText("INVALID CRON EXPRESSION")
			return
		}
		for _, i := range iterations {
			newStr.WriteString(i + "\n")
		}
		hreadableStr.SetText(newStr.String())
	}

	inputField.SetChangedFunc(updateExpr)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
