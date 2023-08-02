package usecase

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// TranslationUseCase 处理业务逻辑
type TranslationUseCase struct {
	repo   TranslationRepo
	webAPI TranslationWebAPI
}

// New 创建业务逻辑
func New(r TranslationRepo, w TranslationWebAPI) *TranslationUseCase {
	return &TranslationUseCase{
		repo:   r,
		webAPI: w,
	}
}

// History 从存储层获取翻译记录
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	// 调用存储层获取翻译记录
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}
	return translations, nil
}

// Translate 翻译
func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	// 调用接口翻译
	translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	// 存储记录
	err = uc.repo.Store(context.Background(), translation)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return translation, nil
}
