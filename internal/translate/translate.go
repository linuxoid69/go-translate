package translate

import (
	"encoding/json"
	"fmt"
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
		return txt, fmt.Errorf("post form with parameters %w", err)
	}

	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)

	if err != nil {
		return txt, fmt.Errorf("read response body %w", err)
	}

	if err:= json.Unmarshal(text, &txt); err != nil {
		return txt, fmt.Errorf("unmarshal error %w", err)
	}

	return txt, nil
}
