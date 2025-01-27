//go:build tview
// +build tview

package renderer

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tview struct{}

func NewRenderer() Renderer {
	return Tview{}
}

func (t Tview) Render(cp Parser) {
	app := tview.NewApplication().EnablePaste(true)

	left := tview.NewFlex().SetDirection(tview.FlexRow)
	left.SetBackgroundColor(tcell.ColorDefault)

	right := tview.NewFlex()
	right.SetBackgroundColor(tcell.ColorDefault)

	separator := tview.NewTextView().SetText(strings.Repeat("│\n", 8))
	separator.SetBackgroundColor(tcell.ColorDefault)
	separator.SetTextColor(tcell.ColorWhite)

	inputField := tview.NewTextArea().
		SetPlaceholder("* * * * *").
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	inputField.SetBackgroundColor(tcell.ColorDefault)

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

	updateExpr := func() {
		hreadableStr.Clear()
		hintsView.Clear()

		text := inputField.GetText()
		text = strings.Trim(text, " ")
		if text == "" {
			return
		}

		splitStr := strings.Split(text, " ")
		inpLen := len(splitStr)

		if inpLen > 5 {
			hreadableStr.SetText("invalid cron expression")
			return
		}

		currElem := splitStr[inpLen-1]
		padding := len(text) - len(currElem)

		hintsView.SetText(cp.GetHints(padding, inpLen-1))

		err := cp.SetExpr(text)
		if err != nil {
			hreadableStr.SetText("invalid cron expression")
			return
		}
		desc, err := cp.HumanReadableStr()
		if err != nil {
			hreadableStr.SetText("invalid cron expression")
			return
		}
		newStr := strings.Builder{}
		newStr.Grow(1024)
		newStr.WriteString(desc + "\n\n")
		newStr.WriteString("Next 3 Iterations:\n\n")
		iterations, err := cp.NextInstances(3)
		if err != nil {
			hreadableStr.SetText("invalid cron expression")
			return
		}
		for _, i := range iterations {
			newStr.WriteString(i + "\n")
		}
		hreadableStr.SetText(newStr.String())
	}

	inputField.SetChangedFunc(updateExpr)
	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		text := inputField.GetText()
		// VALIDATION
		if (text == "" || strings.HasSuffix(text, " ") || len(strings.Split(strings.TrimSuffix(text, " "), " ")) >= 5) && event.Rune() == 32 {
			return nil
		}
		_, fc, _, _ := inputField.GetCursor()
		idx := fc / 2
		if idx > 4 {
			return event
		}
		items := strings.Split(text, " ")
		padding := 0
		for _, item := range items[0:idx] {
			padding += len(item) + 1
		}
		hintsView.SetText(cp.GetHints(padding, idx))
		return event
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
