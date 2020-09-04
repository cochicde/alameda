package node

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/operator/datahub/client"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/pkg/errors"
)

// providerID: aws:///us-west-2a/i-0769ec8570198bf4b --> <provider_raw>//<region>//<instance_id>

// AlamedaNodeRepository creates predicted node to datahub
type AlamedaNodeRepository struct {
	datahubClient *datahubpkg.Client
	clusterUID    string
}

// NewNodeRepository return AlamedaNodeRepository instance
func NewNodeRepository(datahubClient *datahubpkg.Client, clusterUID string) *AlamedaNodeRepository {
	return &AlamedaNodeRepository{
		datahubClient: datahubClient,
		clusterUID:    clusterUID,
	}
}

// CreateNodes creates predicted node to datahub
func (repo *AlamedaNodeRepository) CreateNodes(nodes []entities.ResourceClusterStatusNode) error {
	return repo.datahubClient.Create(&nodes)
}

// DeleteNodes delete predicted node from datahub
func (repo *AlamedaNodeRepository) DeleteNodes(arg interface{}) error {
	objMeta := []*datahub_resources.ObjectMeta{}
	if nodes, ok := arg.([]*datahub_resources.Node); ok {
		for _, node := range nodes {
			objMeta = append(objMeta, (*node).ObjectMeta)
		}
	}
	if meta, ok := arg.([]*datahub_resources.ObjectMeta); ok {
		objMeta = meta
	}
	if nodes, ok := arg.([]entities.ResourceClusterStatusNode); ok {
		for _, node := range nodes {
			objMeta = append(objMeta, &datahub_resources.ObjectMeta{
				Name:        node.Name,
				NodeName:    node.Name,
				ClusterName: node.ClusterName,
				Uid:         node.Uid,
			})
		}
	}

	req := datahub_resources.DeleteNodesRequest{
		ObjectMeta: objMeta,
	}

	if resp, err := repo.datahubClient.DeleteNodes(&req); err != nil {
		return errors.Wrap(err, "delete node from Datahub failed")
	} else if _, err := client.IsResponseStatusOK(resp); err != nil {
		return errors.Wrap(err, "delete nodes from Datahub failed")
	}
	return nil
}

// ListNodes lists nodes to datahub
func (repo *AlamedaNodeRepository) ListNodes() ([]entities.ResourceClusterStatusNode, error) {
	nodes := []entities.ResourceClusterStatusNode{}
	err := repo.datahubClient.List(&nodes, datahubpkg.Option{
		Entity: entities.ResourceClusterStatusNode{
			ClusterName: repo.clusterUID,
		},
		Fields: []string{"ClusterName"},
	})
	if err != nil {
		return nodes, err
	}

	return nodes, nil
}
