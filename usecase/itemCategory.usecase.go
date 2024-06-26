package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type itemCateUsecase struct {
	itemCategoryRepository domain.ItemCategoryRepository
}

func NewItemCategoryUsecase(itemCategoryRepository domain.ItemCategoryRepository) domain.ItemCategoryUseCase {
	return &itemCateUsecase{itemCategoryRepository: itemCategoryRepository}
}

func (u *itemCateUsecase) CreateItemCategory(category *domain.ItemCategory) (*domain.ItemCategory, error) {
	if category.CategoryName == "" {
		return nil, errors.New("erorr! body empty")
	}
	response, err := u.itemCategoryRepository.CreateItemCategory(category)
	if err != nil {
		return nil, errors.New("erorr! cannot create new item category")
	}
	return response, nil
}

func (u *itemCateUsecase) GetAllItemCategory() (*[]domain.ItemCategory, error) {
	response, err := u.itemCategoryRepository.GetAllItemCategory()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemCateUsecase) FindByID(id int) (*domain.ItemCategory, error) {
	response, err := u.itemCategoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemCateUsecase) UpdateItemCategory(id int, category *domain.ItemCategory) (*domain.ItemCategory, error) {
	if category.CategoryName == "" || id == 0 {
		return nil, errors.New("erorr! body empty")
	}
	response, err := u.itemCategoryRepository.UpdateItemCategory(id, category)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *itemCateUsecase) DeleteItemCategory(id int) error {
	err := u.itemCategoryRepository.DeleteItemCategory(id)
	if err != nil {
		return err
	}
	return nil
}
