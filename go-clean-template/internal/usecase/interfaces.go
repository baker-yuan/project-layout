// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation 业务逻辑
	Translation interface {
		// Translate 翻译
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		// History 获取翻译历史
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo 数据存储操作
	TranslationRepo interface {
		// Store 存储翻译记录
		Store(context.Context, entity.Translation) error
		// GetHistory 获取翻译历史
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI 外部接口
	TranslationWebAPI interface {
		// Translate 调用外部接口进行翻译
		Translate(entity.Translation) (entity.Translation, error)
	}
)
