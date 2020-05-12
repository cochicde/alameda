package machinegroup

type ListMachineGroupsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	Name                   string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	Namespace              string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}

type DeleteMachineGroupsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	Name                   string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	Namespace              string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}