package main

type TextWindowState struct {
	text string
}

func (t *TextWindowState) GetText() string {
	return t.text
}

type TextWindow struct {
	currentText string
}

func (t *TextWindow) AddText(str string) {
	t.currentText += str
}

func (t *TextWindow) Save() TextWindowState {
	return TextWindowState{text: t.currentText}
}

func (t *TextWindow) Restore(save TextWindowState) {
	t.currentText = save.GetText()
}

type TextEditor struct {
	textWindow      TextWindow
	savedTextWindow TextWindowState
}

func (t *TextEditor) HitSave() {
	t.savedTextWindow = t.textWindow.Save()
}

func (t *TextEditor) HitUndo() {
	t.textWindow.Restore(t.savedTextWindow)
}
