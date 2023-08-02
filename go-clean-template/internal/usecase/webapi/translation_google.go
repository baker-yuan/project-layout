package webapi

import (
	"fmt"

	translator "github.com/Conight/go-googletrans"
	"github.com/evrone/go-clean-template/internal/entity"
)

// TranslationWebAPI 外部翻译接口
type TranslationWebAPI struct {
	conf translator.Config
}

// New 外部翻译接口
func New() *TranslationWebAPI {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	return &TranslationWebAPI{
		conf: conf,
	}
}

// Translate 调用外部接口翻译
func (t *TranslationWebAPI) Translate(translation entity.Translation) (entity.Translation, error) {
	trans := translator.New(t.conf)

	result, err := trans.Translate(translation.Original, translation.Source, translation.Destination)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	translation.Translation = result.Text

	return translation, nil
}
