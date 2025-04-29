package google_translate

import (
	"testing"

	"github.com/gilang-as/google-translate/params"
	"github.com/gilang-as/google-translate/params/language"
)

func TestTranslateWithParam(t *testing.T) {
	value := params.Translate{
		Text: "Halo Dunia",
		//From: "id",
		To: language.ENGLISH,
	}
	translated, err := TranslateWithParam(value)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}

func TestTranslate(t *testing.T) {
	translated, err := Translate("Hello World", language.INDONESIAN)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}

func TestManualTranslate(t *testing.T) {
	translated, err := ManualTranslate("Halo Semuanya", language.INDONESIAN, language.JAVANESE)
	if err != nil {
		t.Error(err)
	}
	if translated.Text != "" {
		t.Log(translated)
	}
}
