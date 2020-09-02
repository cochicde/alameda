/*
Copyright 2020 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"fmt"

	// apivalidate "github.com/containers-ai/alameda/operator/api/validate"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// ctrl "sigs.k8s.io/controller-runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:validation:Enum=true;false
type enableExecution bool

const (
	defaultEnableExecution enableExecution = false
)

// +kubebuilder:validation:Enum=stable;compact
type alamedaPolicy string

type NamespacedName string

const (
	RecommendationPolicySTABLE  alamedaPolicy = "stable"
	RecommendationPolicyCOMPACT alamedaPolicy = "compact"
)

type AlamedaContainer struct {
	Name      string                      `json:"name" protobuf:"bytes,1,opt,name=name"`
	Resources corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"`
}

type AlamedaPod struct {
	Namespace  string             `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
	Name       string             `json:"name" protobuf:"bytes,2,opt,name=name"`
	UID        string             `json:"uid" protobuf:"bytes,3,opt,name=uid"`
	Containers []AlamedaContainer `json:"containers" protobuf:"bytes,4,opt,name=containers"`
}

func (p *AlamedaPod) GetNamespacedName() NamespacedName {
	return NamespacedName(getNamespacedNameKey(p.Namespace, p.Name))
}

func getNamespacedNameKey(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}

type AlamedaResource struct {
	Namespace    string                        `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
	Name         string                        `json:"name" protobuf:"bytes,2,opt,name=name"`
	UID          string                        `json:"uid" protobuf:"bytes,3,opt,name=uid"`
	Pods         map[NamespacedName]AlamedaPod `json:"pods,omitempty" protobuf:"bytes,4,opt,name=pods"`
	SpecReplicas *int32                        `json:"specReplicas" protobuf:"varint,5,opt,name=spec_replicas"`
	Effective    bool                          `json:"effective" protobuf:"varint,6,opt,name=effective"`
	Message      string                        `json:"message" protobuf:"varint,5,opt,name=message"`
}

func (a AlamedaResource) GetNamespacedName() NamespacedName {
	return NamespacedName(getNamespacedNameKey(a.Namespace, a.Name))
}

type AlamedaController struct {
	Deployments       map[NamespacedName]AlamedaResource `json:"deployments,omitempty" protobuf:"bytes,1,opt,name=deployments"`
	DeploymentConfigs map[NamespacedName]AlamedaResource `json:"deploymentConfigs,omitempty" protobuf:"bytes,2,opt,name=deployment_configs"`
	StatefulSets      map[NamespacedName]AlamedaResource `json:"statefulSets,omitempty" protobuf:"bytes,3,opt,name=stateful_sets"`
}

func NewAlamedaController() AlamedaController {
	return AlamedaController{
		Deployments:       make(map[NamespacedName]AlamedaResource),
		DeploymentConfigs: make(map[NamespacedName]AlamedaResource),
		StatefulSets:      make(map[NamespacedName]AlamedaResource),
	}
}

type AlamedaControllerType int

const (
	DeploymentController       AlamedaControllerType = 1
	DeploymentConfigController AlamedaControllerType = 2
	StatefulSetController      AlamedaControllerType = 3
)

var (
	AlamedaControllerTypeName = map[AlamedaControllerType]string{
		DeploymentController:       "deployment",
		DeploymentConfigController: "deploymentconfig",
		StatefulSetController:      "statefulset",
	}

	K8SKindToAlamedaControllerType = map[string]AlamedaControllerType{
		"Deployment":       DeploymentController,
		"DeploymentConfig": DeploymentConfigController,
		"StatefulSet":      StatefulSetController,
	}
)

type TriggerThreshold struct {
	// +kubebuilder:validation:Pattern=^\d*[1-9]+\d*%$|^\d*[1-9]+\d*\.\d*%$|^\d*\.\d*[1-9]+\d*%$
	CPU string `json:"cpu,omitempty" protobuf:"bytes,1,name=cpu"`
	// +kubebuilder:validation:Pattern=^\d*[1-9]+\d*%$|^\d*[1-9]+\d*\.\d*%$|^\d*\.\d*[1-9]+\d*%$
	Memory string `json:"memory,omitempty" protobuf:"bytes,2,name=memory"`
}

const (
	DefaultTriggerThresholdCPUPercentage    = "10%"
	DefaultTriggerThresholdMemoryPercentage = "10%"
)

func NewDefaultTriggerThreshold() TriggerThreshold {

	return TriggerThreshold{
		CPU:    DefaultTriggerThresholdCPUPercentage,
		Memory: DefaultTriggerThresholdMemoryPercentage,
	}
}

type ExecutionStrategy struct {
	// +kubebuilder:validation:Pattern=^\d*[1-9]+\d*(%?$)$|^\d*[1-9]+\d*\.\d*(%?$)$|^\d*\.\d*[1-9]+\d*(%?$)$
	MaxUnavailable   string                       `json:"maxUnavailable,omitempty" protobuf:"bytes,1,name=max_unavailable"`
	TriggerThreshold *TriggerThreshold            `json:"triggerThreshold,omitempty" protobuf:"bytes,2,name=trigger_threshold"`
	Resources        *corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,3,name=resources"`
}

const (
	DefaultMaxUnavailablePercentage = "25%"
)

func NewDefaultExecutionStrategy() ExecutionStrategy {

	triggerThreshold := NewDefaultTriggerThreshold()

	return ExecutionStrategy{
		MaxUnavailable:   DefaultMaxUnavailablePercentage,
		TriggerThreshold: &triggerThreshold,
	}
}

// +kubebuilder:validation:Enum="";hpa;ca;N/A
type ScalingToolType string

const (
	ScalingToolTypeHPA     = "hpa"
	ScalingToolTypeCA      = "ca"
	ScalingToolTypeDefault = "N/A"
)

type ScalingToolSpec struct {
	Type               ScalingToolType    `json:"type,omitempty" protobuf:"bytes,1,name=type"`
	MinReplicas        *int32             `json:"minReplicas,omitempty" protobuf:"bytes,2,opt,name=min_replicas"`
	MaxReplicas        *int32             `json:"maxReplicas,omitempty" protobuf:"bytes,3,opt,name=max_replicas"`
	ExecutionStrategy  *ExecutionStrategy `json:"executionStrategy,omitempty" protobuf:"bytes,4,name=execution_strategy"`
	CASafeToEvict      bool               `json:"caSafeToEvict,omitempty" protobuf:"bytes,5,name=caSafeToEvict"`
	MachineGroupScaler string             `json:"machineGroupScaler,omitempty" protobuf:"bytes,6,name=machineGroupScaler"`
}

//type AlamedaScalerType = string

const (
	AlamedaScalerTypeNotDefine = ""
	AlamedaScalerTypeDefault   = "default"
	AlamedaScalerTypeKafka     = "kafka"
	AlamedaScalerTypeNginx     = "nginx"
)

type KafkaSpec struct {
	// +kubebuilder:validation:MinLength=1
	ExporterNamespace string `json:"exporterNamespace,omitempty" protobuf:"bytes,1,opt,name=exporter_namespace"`
	// +kubebuilder:validation:MinItems=1
	Topics []string `json:"topics,omitempty" protobuf:"bytes,2,opt,name=topics"`
	// +kubebuilder:validation:MinItems=1
	ConsumerGroups []KafkaConsumerGroupSpec `json:"consumerGroups,omitempty" protobuf:"bytes,3,opt,name=consumer_groups"`
}

type NginxSpec struct {
	Name                    string                `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Selector                *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,2,opt,name=selector"`
	ExporterNamespace       string                `json:"exporterNamespace,omitempty" protobuf:"bytes,3,opt,name=exporter_namespace"`
	ReplicaMarginPercentage *int32                `json:"replicaMarginPercentage,omitempty" protobuf:"bytes,4,opt,name=replica_margin_percentage"`
	MinReplicas             *int32                `json:"minReplicas,omitempty" protobuf:"bytes,5,opt,name=min_replicas"`
	MaxReplicas             *int32                `json:"maxReplicas,omitempty" protobuf:"bytes,6,opt,name=max_replicas"`
	Service                 string                `json:"service,omitempty" protobuf:"bytes,7,opt,name=service"`
	TargetResponseTime      *int64                `json:"targetResponseTime,omitempty" protobuf:"bytes,8,opt,name=target_response_time"`
}

type KafkaConsumerGroupSpec struct {
	// +kubebuilder:validation:MinLength=1
	Name        string                         `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Resource    KafkaConsumerGroupResourceSpec `json:"resource,omitempty" protobuf:"bytes,2,opt,name=resource"`
	MajorTopic  string                         `json:"majorTopic" protobuf:"bytes,3,opt,name=major_topic"`
	MinReplicas *int32                         `json:"minReplicas,omitempty" protobuf:"bytes,4,opt,name=min_replicas"`
	MaxReplicas *int32                         `json:"maxReplicas,omitempty" protobuf:"bytes,5,opt,name=max_replicas"`
}

type KafkaConsumerGroupResourceSpec struct {
	Kubernetes *KubernetesResourceSpec `json:"kubernetes,omitempty" protobuf:"bytes,1,opt,name=kubernetes"`
	Custom     string                  `json:"custom,omitempty" protobuf:"bytes,2,opt,name=custom"`
}

type NginxResourceSpec struct {
	Kubernetes *KubernetesResourceSpec `json:"kubernetes,omitempty" protobuf:"bytes,1,opt,name=kubernetes"`
}

type KubernetesResourceSpec struct {
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,1,opt,name=selector"`
}

// AlamedaScalerSpec defines the desired state of AlamedaScaler
// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
type AlamedaScalerSpec struct {
	Selector              *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,1,name=selector"`
	EnableExecution       *enableExecution      `json:"enableExecution,omitempty" protobuf:"bytes,2,name=enable_execution"`
	Policy                alamedaPolicy         `json:"policy,omitempty" protobuf:"bytes,3,opt,name=policy"`
	CustomResourceVersion string                `json:"customResourceVersion,omitempty" protobuf:"bytes,4,opt,name=custom_resource_version"`
	ScalingTool           ScalingToolSpec       `json:"scalingTool,omitempty" protobuf:"bytes,5,opt,name=scaling_tool"`
	Type                  string                `json:"type,omitempty" protobuf:"bytes,6,opt,name=type"`
	Kafka                 *KafkaSpec            `json:"kafka,omitempty" protobuf:"bytes,7,opt,name=kafka"`
	Nginx                 *NginxSpec            `json:"nginx,omitempty" protobuf:"bytes,8,opt,name=nginx"`
}

type KafkaStatus struct {
	Effective         bool                       `json:"effective" protobuf:"bytes,1,opt,name=effective"`
	Message           string                     `json:"message" protobuf:"bytes,2,opt,name=message"`
	ExporterNamespace string                     `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	Topics            []string                   `json:"topics,omitempty" protobuf:"bytes,4,opt,name=topics"`
	ConsumerGroups    []KafkaConsumerGroupStatus `json:"consumerGroups,omitempty" protobuf:"bytes,5,opt,name=consumer_groups"`
}

type NginxStatus struct {
	Message           string            `json:"message" protobuf:"bytes,1,opt,name=message"`
	AlamedaController AlamedaController `json:"alamedaController,omitempty" protobuf:"bytes,2,opt,name=alameda_controller"`
}

type KafkaConsumerGroupStatus struct {
	Name        string                             `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Topic       string                             `json:"topic,omitempty" protobuf:"bytes,2,opt,name=topic"`
	Resource    KafkaConsumerGroupResourceMetadata `json:"resource,omitempty" protobuf:"bytes,3,opt,name=resource"`
	MinReplicas int32                              `json:"minReplicas,omitempty" protobuf:"bytes,4,opt,name=min_replicas"`
	MaxReplicas int32                              `json:"maxReplicas,omitempty" protobuf:"bytes,5,opt,name=max_replicas"`
}

type KafkaConsumerGroupResourceMetadata struct {
	CustomName string                    `json:"customName,omitempty" protobuf:"bytes,1,opt,name=custom_name"`
	Kubernetes *KubernetesObjectMetadata `json:"kubernetes,omitempty" protobuf:"bytes,2,opt,name=kubernetes"`
}

type KubernetesObjectMetadata struct {
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`
	Name      string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	Kind      string `json:"kind,omitempty" protobuf:"bytes,3,opt,name=kind"`
}

// AlamedaScalerStatus defines the observed state of AlamedaScaler
type AlamedaScalerStatus struct {
	AlamedaController AlamedaController `json:"alamedaController,omitempty" protobuf:"bytes,4,opt,name=alameda_controller"`
	Kafka             *KafkaStatus      `json:"kafka,omitempty" protobuf:"bytes,5,opt,name=kafka"`
	Nginx             *NginxStatus      `json:"nginx,omitempty" protobuf:"bytes,6,opt,name=nginx"`
}

// AlamedaScaler is the Schema for the alamedascalers API
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Scaling Type",type=string,JSONPath=`.spec.type`,description="The type of scaling"
// +kubebuilder:printcolumn:name="Scaling Tool",type=string,JSONPath=`.spec.scalingTool.type`,description="The type of scaling tool"
// +kubebuilder:printcolumn:name="Execution",type=boolean,JSONPath=`.spec.enableExecution`,description="The enable of execution"
// +kubebuilder:printcolumn:name="Selector",type=string,JSONPath=`.spec.selector.matchLabels`,description="The match labels of selector"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="The time of creation"
type AlamedaScaler struct {
	// Mgr      ctrl.Manager                      `json:"-"`
	// Validate apivalidate.AlamedaScalerValidate `json:"-"`

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaScalerSpec   `json:"spec,omitempty"`
	Status AlamedaScalerStatus `json:"status,omitempty"`
}

func (as *AlamedaScaler) SetDefaultValue() { //this function is set alamedascaler default value
	as.setDefaultEnableExecution()
	as.setDefaultScalingTool()
}

func (as *AlamedaScaler) SetCustomResourceVersion(v string) {
	as.Spec.CustomResourceVersion = v
}

func (as *AlamedaScaler) SetStatusAlamedaController(ac AlamedaController) {
	as.Status.AlamedaController = ac
}

func (as *AlamedaScaler) SetStatusKafka(k *KafkaStatus) {
	as.Status.Kafka = k
}

func (as *AlamedaScaler) SetType(t string) {
	as.Spec.Type = t
}

func (as *AlamedaScaler) GenCustomResourceVersion() string {
	v := as.ResourceVersion
	return v
}

func (as *AlamedaScaler) GetType() string {
	return as.Spec.Type
}

func (as *AlamedaScaler) GetKafkaNamespace() string {

	if as.Spec.Kafka == nil {
		return ""
	}

	return as.Spec.Kafka.ExporterNamespace
}

func (as *AlamedaScaler) ListKafkaTopics() []string {

	if as.Spec.Kafka == nil {
		return nil
	}

	return as.Spec.Kafka.Topics
}

func (as *AlamedaScaler) ListKafkaConsumerGroupSpecs() []KafkaConsumerGroupSpec {

	if as.Spec.Kafka == nil {
		return nil
	}

	return as.Spec.Kafka.ConsumerGroups
}

// GetMonitoredPods returns pods restoring in AlamedaScaler.Status
func (as *AlamedaScaler) GetMonitoredPods() []*AlamedaPod {
	pods := make([]*AlamedaPod, 0)

	for _, alamedaResource := range as.Status.AlamedaController.Deployments {
		for _, pod := range alamedaResource.Pods {
			cpPod := pod
			pods = append(pods, &cpPod)
		}
	}

	for _, alamedaResource := range as.Status.AlamedaController.DeploymentConfigs {
		for _, pod := range alamedaResource.Pods {
			cpPod := pod
			pods = append(pods, &cpPod)
		}
	}

	for _, alamedaResource := range as.Status.AlamedaController.StatefulSets {
		for _, pod := range alamedaResource.Pods {
			cpPod := pod
			pods = append(pods, &cpPod)
		}
	}

	return pods
}

func (as *AlamedaScaler) IsEnableExecution() bool {
	if as.Spec.EnableExecution == nil ||
		*as.Spec.EnableExecution == false {
		return false
	}
	return true
}

func (as *AlamedaScaler) IsScalingToolTypeHPA() bool {
	return as.Spec.ScalingTool.Type == ScalingToolTypeHPA
}

func (as *AlamedaScaler) IsScalingToolTypeCA() bool {
	return as.Spec.ScalingTool.Type == ScalingToolTypeCA
}

// HasAlamedaPod returns true if the pod is reocording in AlamedaScaler.Status
func (as *AlamedaScaler) HasAlamedaPod(namespace, name string) bool {

	for _, deployment := range as.Status.AlamedaController.Deployments {
		deploymentNS := deployment.Namespace
		for _, pod := range deployment.Pods {
			if deploymentNS == namespace && pod.Name == name {
				return true
			}
		}
	}
	for _, deploymentConfig := range as.Status.AlamedaController.DeploymentConfigs {
		deploymentConfigNS := deploymentConfig.Namespace
		for _, pod := range deploymentConfig.Pods {
			if deploymentConfigNS == namespace && pod.Name == name {
				return true
			}
		}
	}
	for _, statefulSet := range as.Status.AlamedaController.StatefulSets {
		statefulSetNS := statefulSet.Namespace
		for _, pod := range statefulSet.Pods {
			if statefulSetNS == namespace && pod.Name == name {
				return true
			}
		}
	}

	return false
}

func (as *AlamedaScaler) AddAlamedaResourceIntoStatus(
	arType AlamedaControllerType, ar AlamedaResource) {
	ac := as.Status.AlamedaController
	switch arType {
	case DeploymentController:
		ac.Deployments[ar.GetNamespacedName()] = ar
	case DeploymentConfigController:
		ac.DeploymentConfigs[ar.GetNamespacedName()] = ar
	case StatefulSetController:
		ac.StatefulSets[ar.GetNamespacedName()] = ar
	}
}

func (as *AlamedaScaler) setDefaultEnableExecution() {
	if as.Spec.EnableExecution == nil {
		copyDefaultEnableExecution := defaultEnableExecution
		as.Spec.EnableExecution = &copyDefaultEnableExecution
	}
}

func (as *AlamedaScaler) setDefaultScalingTool() {

	if as.Spec.ScalingTool.Type == "" {
		as.Spec.ScalingTool.Type = ScalingToolTypeDefault
	}
}

func (as *AlamedaScaler) setDefaultTriggerThreshold() {
	defaultTriggerThreshold := NewDefaultTriggerThreshold()
	as.Spec.ScalingTool.ExecutionStrategy.TriggerThreshold =
		&defaultTriggerThreshold
}

// +kubebuilder:object:root=true
// AlamedaScalerList contains a list of AlamedaScaler
type AlamedaScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaScaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaScaler{}, &AlamedaScalerList{})
}
