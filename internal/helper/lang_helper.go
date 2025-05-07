package helper

import (
	"errors"

	"github.com/pemistahl/lingua-go"
)

type LangHelper struct{}

func (l *LangHelper) CheckLang(text string) error {
	detector := lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithout(lingua.Japanese).Build()
	if _, exists := detector.DetectLanguageOf(text); exists {
		return errors.New("language not supported")
	}
	return nil
}
