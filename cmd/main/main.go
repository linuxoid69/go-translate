package main

import (
	"fmt"
	"os"

	noti "github.com/linuxoid69/go-translate/internal/notification"
	"github.com/linuxoid69/go-translate/internal/translate"
)

func main() {
	var t translate.Translate
	t.URL = "https://translate.mentality.rip/translate"
	t.Params.QueryText = os.Args[1:]
	t.Params.SourceLanguage = "en"
	t.Params.TargetLanguage = "ru"

	lenghWords := len(os.Args[1:])
	if lenghWords != 0 {
		text, err := t.GetTranslate()
		if err != nil {
			noti.Notify(fmt.Sprintf("Error: %v", err))
		}

		noti.Notify(text.Text)
	}

}
