package translate

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Translate //
type Translate struct {
	URL    string
	Params struct {
		QueryText      []string
		SourceLanguage string
		TargetLanguage string
		Format         string
	}
}

// ResponseText //
type ResponseText struct {
	Text string `json:"translatedText"`
}

// GetTranslate - get translate text
func (t *Translate) GetTranslate() (txt ResponseText, err error) {

	res, err := http.PostForm(t.URL,
		url.Values{"q": {strings.Join(t.Params.QueryText, " ")},
			"source": {t.Params.SourceLanguage},
			"target": {t.Params.TargetLanguage},
			"format": {t.Params.Format}})

	if err != nil {
		return txt, err
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)

	json.Unmarshal(text, &txt)

	if err != nil {
		return txt, err
	}

	return txt, nil
}
