package influxdb

import (
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	RepoInfluxPrediction "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/predictions"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

type NamespacePredictions struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNamespacePredictionsWithConfig(config InfluxDB.Config) DaoPredictionTypes.NamespacePredictionsDAO {
	return &NamespacePredictions{InfluxDBConfig: config}
}

// CreateNamespacePredictions Implementation of prediction dao interface
func (p *NamespacePredictions) CreatePredictions(predictions DaoPredictionTypes.NamespacePredictionMap) error {
	predictionRepo := RepoInfluxPrediction.NewNamespaceRepositoryWithConfig(p.InfluxDBConfig)

	err := predictionRepo.CreatePredictions(predictions)
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *NamespacePredictions) ListPredictions(request DaoPredictionTypes.ListNamespacePredictionsRequest) (DaoPredictionTypes.NamespacePredictionMap, error) {
	namespacePredictionMap := DaoPredictionTypes.NewNamespacePredictionMap()

	predictionRepo := RepoInfluxPrediction.NewNamespaceRepositoryWithConfig(p.InfluxDBConfig)
	namespacePredictions, err := predictionRepo.ListPredictions(request)
	if err != nil {
		scope.Error(err.Error())
		return DaoPredictionTypes.NewNamespacePredictionMap(), err
	}
	for _, namespacePrediction := range namespacePredictions {
		namespacePredictionMap.AddNamespacePrediction(namespacePrediction)
	}

	return namespacePredictionMap, nil
}
