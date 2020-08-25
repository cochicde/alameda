package prometheus

import (
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
)

type nodeMetricsFetchingFunction func(nodeName string, options ...DBCommon.Option) ([]Entity, error)
