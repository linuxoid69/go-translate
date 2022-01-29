package main

import (
	"fmt"

	"github.com/linuxoid69/go-translate/internal/clipboard"
	noti "github.com/linuxoid69/go-translate/internal/notification"
	"github.com/linuxoid69/go-translate/internal/translate"
)

func main() {
	var t translate.Translate

	txt, err := clipboard.GetTextFromClipboard()
	if err != nil {
		noti.Notify(fmt.Sprintf("%v", err))
	}

	t.URL = "https://trans.zillyhuhn.com/translate"
	t.Params.QueryText = []string{txt}
	t.Params.SourceLanguage = "auto"
	t.Params.TargetLanguage = "ru"

	text, err := t.GetTranslate()
	if err != nil {
		string("Error: %v", err)
	}

	noti.Notify(text.Text)
}
