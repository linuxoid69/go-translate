package main

import (
	"log"
	"github.com/linuxoid69/go-translate/internal/translate"
	noti "github.com/linuxoid69/go-translate/internal/notification"
	"os"

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
			log.Printf("Can't get text: %v", err)
		}

		noti.Notify(text.Text)
	}


}
