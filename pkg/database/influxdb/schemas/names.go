package schemas

var MeasurementNameMap = map[Scope]string{
	Application:    "application",
	Config:         "config",
	Execution:      "execution",
	Fedemeter:      "fedemeter",
	Metric:         "metric",
	Planning:       "planning",
	Prediction:     "prediction",
	Recommendation: "recommendation",
	Resource:       "resource",
	Target:         "target",
}

var MeasurementSchemaNameMap = map[Scope]string{
	Application:    "application_schema",
	Config:         "config_schema",
	Execution:      "execution_schema",
	Fedemeter:      "fedemeter_schema",
	Metric:         "metric_schema",
	Planning:       "planning_schema",
	Prediction:     "prediction_schema",
	Recommendation: "recommendation_schema",
	Resource:       "resource_schema",
	Target:         "target_schema",
}

var DatabaseNameMap = map[Scope]string{
	Application:    "alameda_application",
	Config:         "alameda_config",
	Execution:      "alameda_execution",
	Fedemeter:      "alameda_fedemeter",
	Metric:         "alameda_metric",
	Planning:       "alameda_planning",
	Prediction:     "alameda_prediction",
	Recommendation: "alameda_recommendation",
	Resource:       "alameda_cluster_status",
	Target:         "alameda_target",
}
