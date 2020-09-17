package tenant

import (
	"encoding/json"
	"strings"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	tenantv1alpha1 "github.com/containers-ai/alameda/operator/apis/tenant/v1alpha1"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
)

func CreateOrganization(datahubClient *datahubpkg.Client,
	org *tenantv1alpha1.AlamedaOrganization) error {
	if org == nil {
		return nil
	}
	err := datahubClient.Create(&[]entities.ConfigTenancyOrganization{
		{
			Name:   org.GetName(),
			Tenant: org.Spec.Tenant,
		},
	})
	if err != nil {
		return err
	}

	tenantClusterEntities := []entities.ConfigTenancyCluster{}
	for _, cluster := range org.Spec.Clusters {
		namespaces := []string{}
		for _, nsName := range cluster.WatchedNamespace.Names {
			namespaces = append(namespaces, nsName)
		}
		keys := []string{}
		for _, key := range cluster.DataSource.Keys {
			keys = append(keys, key.Key)
		}
		tenantClusterEntity := entities.ConfigTenancyCluster{
			Name:                 cluster.Name,
			Organization:         org.GetName(),
			Tenant:               org.Spec.Tenant,
			DataSource:           entities.DataSource(cluster.DataSource.Type),
			DataSourceAddress:    cluster.DataSource.Address,
			DataSourceAccount:    cluster.DataSource.Account,
			WatchedNamespaces:    strings.Join(namespaces, ","),
			DataSourceKeys:       strings.Join(keys, ","),
			ResourcePlanning:     true,                                                            // default to true
			ResourcePlanningMode: entities.DataStoredMode(tenantv1alpha1.UploadResultFeatureMode), // default to upload
			CostAnalysis:         true,                                                            // default to true
			CostAnalysisMode:     entities.DataStoredMode(tenantv1alpha1.UploadResultFeatureMode), // default to upload
		}
		if cluster.WatchedNamespace.Operator != "" {
			tenantClusterEntity.WatchedNamespacesOperator =
				entities.LogicOperator(cluster.WatchedNamespace.Operator)
		}

		if org.Spec.ResourcePlanning != nil {
			tenantClusterEntity.ResourcePlanning = bool(org.Spec.ResourcePlanning.Enabled)
			tenantClusterEntity.ResourcePlanningMode = entities.DataStoredMode(org.Spec.ResourcePlanning.Mode)
		}
		if org.Spec.CostAnalysis != nil {
			tenantClusterEntity.CostAnalysis = bool(org.Spec.CostAnalysis.Enabled)
			tenantClusterEntity.CostAnalysisMode = entities.DataStoredMode(org.Spec.CostAnalysis.Mode)
		}
		for _, feature := range cluster.Features {
			if feature.Type == tenantv1alpha1.ResourcePlanningFeatureType && feature.ResourcePlanning != nil {
				tenantClusterEntity.ResourcePlanning = bool(feature.ResourcePlanning.Enabled)
				tenantClusterEntity.ResourcePlanningMode = entities.DataStoredMode(feature.ResourcePlanning.Mode)
			} else if feature.Type == tenantv1alpha1.CostAnalysisFeatureType && feature.CostAnalysis != nil {
				tenantClusterEntity.CostAnalysis = bool(feature.CostAnalysis.Enabled)
				tenantClusterEntity.CostAnalysisMode = entities.DataStoredMode(feature.CostAnalysis.Mode)
			}
		}

		clusterRawSpec, err := json.Marshal(cluster)
		if err == nil {
			tenantClusterEntity.RawSpec = string(clusterRawSpec)
		}
		tenantClusterEntities = append(tenantClusterEntities, tenantClusterEntity)
	}
	return datahubClient.Create(&tenantClusterEntities)
}
