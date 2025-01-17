/*
Copyright 2025 SAP SE.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// ValkeySpec defines the desired state of Valkey.
type ValkeySpec struct {
	// Uncomment the following if you want to implement the PlacementConfiguration interface
	// (that is, want to make deployment namespace and name configurable here in the spec, independently of
	// the component's metadata.namespace and metadata.name).
	// component.PlacementSpec                  `json:",inline"`
	// Uncomment the following if you want to implement the ClientConfiguration interface
	// (that is, want to allow remote deployments via a specified kubeconfig).
	// Note, that when implementing the ClientConfiguration interface, then also the PlacementConfiguration
	// interface should be implemented.
	// component.ClientSpec        `json:",inline"`
	// Uncomment the following if you want to implement the ImpersonationConfiguratio interface
	// (that is, want to allow use a specified service account in the target namespace for the deployment).
	// component.ImpersonationSpec `json:",inline"`
	// Uncomment the following if you want to implement the RequeueConfiguration interface
	// (that is, want to allow to override the default requeue interval of 10m).
	// component.RequeueSpec `json:",inline"`
	// Uncomment the following if you want to implement the RetryConfiguration interface
	// (that is, want to allow to override the default retry interval, which equals the effective requeue interval).
	// component.RetrySpec `json:",inline"`

	// Add your own fields here, describing the deployment of the managed component.
	Replicas int                 `json:"replicas,omitempty"`
	Metrics  *MetricsProperties  `json:"metrics,omitempty"`
	Sentinel *SentinelProperties `json:"sentinel,omitempty"`
}

// MetricsProperties defines the properties for metrics configuration.
type MetricsProperties struct {
	Enabled                                 bool `json:"enabled,omitempty"`
	component.KubernetesContainerProperties `json:",inline"`
	ServiceMonitor                          *MetricsServiceMonitorProperties `json:"monitor,omitempty"`
	PrometheusRule                          *MetricsPrometheusRuleProperties `json:"prometheusRule,omitempty"`
}

// SentinelProperties defines the properties for sentinel configuration.
type SentinelProperties struct {
	Enabled bool `json:"enabled,omitempty"`
}

// MetricsPrometheusRuleProperties defines the properties for Prometheus rule configuration.
type MetricsPrometheusRuleProperties struct {
	Enabled bool `json:"enabled,omitempty"`
}

// MetricsServiceMonitorProperties defines the properties for service monitor configuration.
type MetricsServiceMonitorProperties struct {
	Enabled bool `json:"enabled,omitempty"`
}

// ValkeyStatus defines the observed state of Valkey.
type ValkeyStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// Valkey is the Schema for the valkeies API.
type Valkey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ValkeySpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status ValkeyStatus `json:"status,omitempty"`
}

var _ component.Component = &Valkey{}

// +kubebuilder:object:root=true

// ValkeyList contains a list of Valkey.
type ValkeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Valkey `json:"items"`
}

func (s *ValkeySpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *Valkey) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *Valkey) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&Valkey{}, &ValkeyList{})
}
