package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoCluster "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	resources2 "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/resources"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiResources "prophetstor.com/api/datahub/resources"
)

func (s *ServiceV1alpha1) CreateControllers(ctx context.Context, in *ApiResources.CreateControllersRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateControllers grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetControllers() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	requestExtended := resources.CreateControllersRequestExtended{CreateControllersRequest: in}
	if requestExtended.Validate() != nil {
		return &status.Status{
			Code: int32(code.Code_INVALID_ARGUMENT),
		}, nil
	}

	controllerDAO := DaoCluster.NewControllerDAO(*s.Config)
	if err := controllerDAO.CreateControllers(requestExtended.ProduceControllers()); err != nil {
		scope.Errorf("failed to create controllers: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListControllers(ctx context.Context, in *ApiResources.ListControllersRequest) (*ApiResources.ListControllersResponse, error) {
	scope.Debug("Request received from ListControllers grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.ListControllersRequestExtended{ListControllersRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiResources.ListControllersResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	controllerDAO := DaoCluster.NewControllerDAO(*s.Config)
	ctls, err := controllerDAO.ListControllers(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListControllers failed: %+v", err)
		return &ApiResources.ListControllersResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	controllers := make([]*ApiResources.Controller, 0)
	for _, ctl := range ctls {
		controllerExtended := resources2.ControllerExtended{Controller: ctl}
		controller := controllerExtended.ProduceController()
		controllers = append(controllers, controller)
	}

	response := ApiResources.ListControllersResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Controllers: controllers,
	}

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteControllers(ctx context.Context, in *ApiResources.DeleteControllersRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteControllers grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.DeleteControllersRequestExtended{DeleteControllersRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	controllerDAO := DaoCluster.NewControllerDAO(*s.Config)
	if err := controllerDAO.DeleteControllers(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete controllers: %+v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
