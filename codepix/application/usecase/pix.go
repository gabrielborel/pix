package usecase

import (
	"errors"

	"github.com/gabrielborel/pix/codepix/domain/model"
)

type PixKeyUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func NewPixKeyUseCase(repository model.PixKeyRepositoryInterface) *PixKeyUseCase {
	return &PixKeyUseCase{PixKeyRepository: repository}
}

func (uc *PixKeyUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	account, err := uc.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	uc.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

func (uc *PixKeyUseCase) FindKey(key, kind string) (*model.PixKey, error) {
	pixKey, err := uc.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
