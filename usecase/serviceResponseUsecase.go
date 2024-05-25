package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
)

type serviceResponseUsecase struct {
	responseRepo     domain.ServiceResponseRepository
	user             domain.UserUseCase
	requestedService domain.ServiceRequestUsecase
	item             domain.ItemUseCase
	itemRepo         domain.ItemRepository
	itemLog          domain.ItemLogUsecase
	MParts           domain.MaintenancePartsRepository
}

func NewServiceResponsUsecase(responseRepo domain.ServiceResponseRepository,
	user domain.UserUseCase, requestedService domain.ServiceRequestUsecase,
	item domain.ItemUseCase, itemRepo domain.ItemRepository, MParts domain.MaintenancePartsRepository,
	itemLog domain.ItemLogUsecase) domain.ServiceResponseUsecase {
	return &serviceResponseUsecase{
		responseRepo:     responseRepo,
		user:             user,
		requestedService: requestedService,
		item:             item,
		itemRepo:         itemRepo,
		MParts:           MParts,
		itemLog:          itemLog,
	}
}

func (u *serviceResponseUsecase) GetAllResponse() (*[]domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetAllResponse()
	return response, err
}

func (u *serviceResponseUsecase) CreateServiceResponse(newResponse *domain.ServiceResponse) (*domain.ServiceResponse, error) {
	_, userErr := u.user.GetById(newResponse.StaffID)
	if userErr != nil {
		return nil, errors.New("serviceResponse: staffID error")
	}
	_, serviceErr := u.requestedService.GetServiceRequest(newResponse.RequestedServiceID)
	if serviceErr != nil {
		return nil, errors.New("serviceResponnse: not a valid service request")
	}

	for _, v := range newResponse.MaintenanceParts {
		item, err := u.item.FindByID(v.ItemID)
		if err != nil || item.ItemQty-v.ItemQty < 0 {
			return nil, errors.New("item : not valid or out of stock")
		}
	}

	response, responseErr := u.responseRepo.CreateServiceResponse(newResponse)
	if responseErr != nil {
		return nil, responseErr
	}

	for _, v := range newResponse.MaintenanceParts {
		err := u.MParts.CreateMaintenanceParts(response.StaffServiceID, v.ItemID, v.ItemQty, newResponse.CreatedDate)
		if err != nil {
			return nil, err
		}

		newLog := new(domain.ItemLog)
		newLog.ItemID = v.ItemID
		newLog.ItemQty = v.ItemQty
		newLog.StaffID = newResponse.StaffID
		newLog.IsAdd = false
		newLog.CreateDate = newResponse.CreatedDate
		u.itemLog.CreateItemLog(newLog)
	}
	return response, nil
}

func (u *serviceResponseUsecase) DeleteResponse(id int) error {
	return u.responseRepo.DeleteResponse(id)
}

func (u *serviceResponseUsecase) GetResponse(id int) (*domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetResponse(id)
	if err != nil {
		return nil, errors.New("serviceResponse: responseID error")
	}
	return response, err
}

func (u *serviceResponseUsecase) GetResponseByService(id int) (*[]domain.ServiceResponse, error) {
	response, err := u.responseRepo.GetResponseByService(id)
	if err != nil {
		return nil, errors.New("serviceResponse: serviceID error")
	}
	return response, err
}
