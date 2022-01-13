package translate

import (
	"reflect"
	"testing"
)

func TestTranslate_GetTranslate(t *testing.T) {
	type fields struct {
		URL    string
		Params struct {
			QueryText      []string
			SourceLanguage string
			TargetLanguage string
			Format         string
		}
	}
	tests := []struct {
		name    string
		fields  fields
		wantTxt ResponseText
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Translate{
				URL:    tt.fields.URL,
				Params: tt.fields.Params,
			}
			gotTxt, err := tr.GetTranslate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate.GetTranslate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTxt, tt.wantTxt) {
				t.Errorf("Translate.GetTranslate() = %v, want %v", gotTxt, tt.wantTxt)
			}
		})
	}
}
