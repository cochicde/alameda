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

// CreatePods add containers information of pods to database
func (s *ServiceV1alpha1) CreatePods(ctx context.Context, in *ApiResources.CreatePodsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreatePods grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetPods() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	requestExtended := resources.CreatePodsRequestExtended{CreatePodsRequest: in}
	if err := requestExtended.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	podDAO := DaoCluster.NewPodDAO(*s.Config)
	if err := podDAO.CreatePods(requestExtended.ProducePods()); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

// ListAlamedaPods returns predicted pods
func (s *ServiceV1alpha1) ListPods(ctx context.Context, in *ApiResources.ListPodsRequest) (*ApiResources.ListPodsResponse, error) {
	scope.Debug("Request received from ListAlamedaPods grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.ListPodsRequestExtended{ListPodsRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiResources.ListPodsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	podDAO := DaoCluster.NewPodDAO(*s.Config)
	pds, err := podDAO.ListPods(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListNodes failed: %+v", err)
		return &ApiResources.ListPodsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	pods := make([]*ApiResources.Pod, 0)
	for _, pd := range pds {
		podExtended := resources2.PodExtended{Pod: pd}
		pod := podExtended.ProducePod()
		pods = append(pods, pod)
	}

	response := ApiResources.ListPodsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Pods: pods,
	}

	return &response, nil
}

// DeletePods update containers information of pods to database
func (s *ServiceV1alpha1) DeletePods(ctx context.Context, in *ApiResources.DeletePodsRequest) (*status.Status, error) {
	scope.Debug("Request received from DeletePods grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.DeletePodsRequestExtended{DeletePodsRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	podDAO := DaoCluster.NewPodDAO(*s.Config)
	if err := podDAO.DeletePods(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete pods: %+v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
