package handlers

import (
	services "service-product/services/product"
)

type ServiceProductHandler struct {
	serviceResult services.ServiceResult
	serviceResults services.ServiceResults
	serviceCreate services.ServiceCreate
	serviceUpdate services.ServiceUpdate
	serviceDelete services.ServiceDelete
}

func NewHandlerRPCProduct(
	serviceResults services.ServiceResults,
	serviceResult services.ServiceResult,
	serviceCreate services.ServiceCreate,
	serviceUpdate services.ServiceUpdate,
	serviceDelete services.ServiceDelete) *ServiceProductHandler {
	return &ServiceProductHandler{
		serviceResult: serviceResult,
		serviceResults: serviceResults,
		serviceCreate: serviceCreate,
		serviceUpdate: serviceUpdate,
		serviceDelete: serviceDelete,
	}
}





