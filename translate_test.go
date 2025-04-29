package google_translate

import (
	"testing"

	"github.com/is3ka1/google-translate/params"
	"github.com/is3ka1/google-translate/params/language"
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

func TestTranslateLongText(t *testing.T) {
	// 4999 characters of long text
	longText := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vitae consequat dolor. Curabitur tincidunt dolor sit amet euismod lacinia. Sed nec fermentum elit. Vestibulum sagittis ante justo, eget venenatis est laoreet a. Phasellus cursus semper diam, sed finibus nulla tempus vitae. In hac habitasse platea dictumst. Nullam nec cursus neque, ut suscipit eros. Aenean tortor ligula, pulvinar vel tristique in, ornare et tellus. Duis rutrum quam enim, aliquet posuere lacus rutrum at. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae;

Integer a elit mi. Maecenas quis nisl lacinia, vulputate ipsum eu, sagittis erat. Nullam iaculis erat ac neque condimentum accumsan. Ut ultrices metus eget ante sagittis iaculis. Maecenas tortor nunc, mattis quis enim ac, tempus pharetra tortor. Aliquam venenatis eleifend efficitur. Fusce tristique lectus eu leo tempor, sit amet vehicula erat volutpat. Aenean blandit efficitur risus ut lacinia. Curabitur scelerisque cursus dolor, in laoreet velit sagittis ut. Maecenas tempor, risus at volutpat ullamcorper, sem massa luctus lacus, eu faucibus erat libero nec ex. Vivamus sollicitudin ex at erat aliquam elementum. Donec pellentesque ante quis interdum iaculis. In pellentesque viverra sem, ac tempor metus hendrerit sed. Cras ac blandit velit, vitae auctor neque.

Sed ullamcorper condimentum libero eget commodo. Nulla efficitur dictum nibh, sed elementum nibh consequat quis. Donec vel risus eget sapien feugiat commodo. Proin tincidunt neque at ipsum pulvinar porta. Ut porta molestie sem. Mauris iaculis dolor quis cursus molestie. Donec porta lorem sit amet nisl sollicitudin, eget faucibus tellus iaculis. Ut congue venenatis purus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Fusce non diam tristique, laoreet tortor a, feugiat nulla. Curabitur ut sem augue. Proin sagittis finibus lectus. Quisque cursus lacus id urna aliquet, quis dignissim sem sagittis. Nulla risus mauris, pharetra pretium odio at, interdum dignissim lacus. Nunc faucibus mauris eu cursus vehicula. Nunc et porta neque.

Aliquam id maximus nisl. Quisque at massa in nibh feugiat accumsan. Nunc augue felis, viverra a nisi a, blandit placerat justo. Quisque diam velit, mattis nec neque eu, mattis rutrum ex. Nam tempus laoreet malesuada. Pellentesque pharetra ac massa sit amet euismod. Curabitur rutrum elementum bibendum. Suspendisse potenti. Integer nec tincidunt libero, sit amet commodo turpis. Nunc tristique, neque et blandit dignissim, leo arcu scelerisque felis, vel pretium metus lorem et nunc. Aliquam fringilla condimentum augue. Nulla convallis vel velit at facilisis.

Quisque at pellentesque augue, nec convallis nisl. Fusce hendrerit fermentum nibh, vel sollicitudin purus rutrum tristique. Aliquam sed quam turpis. Etiam mattis tristique finibus. Sed consequat semper dolor, vel eleifend sapien dapibus vel. Vestibulum elit erat, egestas non pellentesque pulvinar, lacinia in velit. Praesent egestas arcu et ullamcorper tincidunt. Nam consectetur leo eget elit mollis venenatis. Donec iaculis ultrices posuere. In at ultrices sapien. Nullam at maximus arcu.

Aenean nec lectus nulla. Ut arcu purus, rutrum at mattis ut, porta in leo. Nulla metus velit, blandit ut tincidunt id, viverra vitae nisi. Vivamus malesuada est ex, feugiat sodales velit pulvinar in. Morbi porta nunc non dignissim ullamcorper. Integer erat odio, imperdiet ac fringilla ac, aliquam ut turpis. Ut vulputate molestie viverra. Duis faucibus eros at imperdiet ornare.

Etiam quis dignissim nunc. Vestibulum tincidunt nulla vel leo egestas cursus. Sed ligula libero, placerat sed gravida et, luctus ac sapien. Suspendisse potenti. Cras ultricies varius lacus, et maximus libero gravida quis. Vivamus id leo vel justo tristique tempor. Quisque massa mauris, rutrum vel leo vehicula, rutrum pharetra elit. Integer vulputate vel purus id condimentum. Donec tristique porta tortor, id consectetur leo volutpat et. Nulla in ante non neque maximus sollicitudin. Mauris ut dignissim tortor, sit amet vestibulum nisl. Donec imperdiet, magna nec viverra semper, lectus nisl facilisis lacus, et congue ligula arcu quis neque. Phasellus posuere metus elementum dolor pulvinar sodales. Donec in elit justo. Nulla laoreet aliquet augue, non suscipit orci mollis sit amet.

Nunc sed dui mattis, faucibus velit non, cursus dolor. Morbi luctus risus lacus, sit amet semper nulla eleifend non. Ut rutrum pharetra justo, eget malesuada libero semper ac. Sed laoreet nulla a lacus rutrum convallis. Proin sodales quis nibh eget fringilla. Nulla magna tellus, aliquam id lacinia id, dictum a magna. Nam placerat sem velit, vitae semper velit commodo id. Pellentesque erat tortor, tristique ut convallis vel, scelerisque sed urna. Mauris dictum egestas nisl, in rutrum elit tempor in. In egestas lorem nec ante fermentum aliquet. Sed mauris orci, tincidunt vitae est eget, interdum posuere sapien. Proin ut tempus ipsum. Aliquam urna.`
	translated, err := Translate(longText, language.ENGLISH)
	if err != nil {
		t.Error(err)
		return // Stop test if translation failed
	}
	t.Logf("Length of original text: %d", len(longText))          // Log original length
	t.Logf("Length of translated text: %d", len(translated.Text)) // Log the length

	// Check if translated length is at least 80% of original length
	minLength := float64(len(longText)) * 0.80
	if float64(len(translated.Text)) < minLength {
		t.Errorf("Translated text length (%d) is less than 80%% of original length (%d, threshold: %.0f)",
			len(translated.Text), len(longText), minLength)
	}

	if translated.Text != "" { // Keep logging the actual result
		t.Log(translated)
	}
}
