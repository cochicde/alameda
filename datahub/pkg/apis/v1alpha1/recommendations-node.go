package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoRecommendation "prophetstor.com/alameda/datahub/pkg/dao/interfaces/recommendations"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

func (s *ServiceV1alpha1) CreateNodeRecommendations(ctx context.Context, in *ApiRecommendations.CreateNodeRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateNodeRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	nodeRecommendationList := in.GetNodeRecommendations()
	nodeDAO := DaoRecommendation.NewNodeRecommendationsDAO(*s.Config)
	err := nodeDAO.CreateRecommendations(nodeRecommendationList)

	if err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListNodeRecommendations(ctx context.Context, in *ApiRecommendations.ListNodeRecommendationsRequest) (*ApiRecommendations.ListNodeRecommendationsResponse, error) {
	scope.Debug("Request received from ListNodeRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	nodeDAO := DaoRecommendation.NewNodeRecommendationsDAO(*s.Config)
	nodeRecommendations, err := nodeDAO.ListRecommendations(in)
	if err != nil {
		scope.Errorf("api ListNodeRecommendations failed: %v", err)
		response := &ApiRecommendations.ListNodeRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
			NodeRecommendations: nodeRecommendations,
		}
		return response, nil
	}

	response := &ApiRecommendations.ListNodeRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		NodeRecommendations: nodeRecommendations,
	}

	return response, nil
}
