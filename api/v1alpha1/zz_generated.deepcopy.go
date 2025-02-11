//go:build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and valkey-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingProperties) DeepCopyInto(out *BindingProperties) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingProperties.
func (in *BindingProperties) DeepCopy() *BindingProperties {
	if in == nil {
		return nil
	}
	out := new(BindingProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerProperties) DeepCopyInto(out *CertManagerProperties) {
	*out = *in
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerProperties.
func (in *CertManagerProperties) DeepCopy() *CertManagerProperties {
	if in == nil {
		return nil
	}
	out := new(CertManagerProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsPrometheusRuleProperties) DeepCopyInto(out *MetricsPrometheusRuleProperties) {
	*out = *in
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]monitoringv1.Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsPrometheusRuleProperties.
func (in *MetricsPrometheusRuleProperties) DeepCopy() *MetricsPrometheusRuleProperties {
	if in == nil {
		return nil
	}
	out := new(MetricsPrometheusRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsProperties) DeepCopyInto(out *MetricsProperties) {
	*out = *in
	in.KubernetesContainerProperties.DeepCopyInto(&out.KubernetesContainerProperties)
	if in.ServiceMonitor != nil {
		in, out := &in.ServiceMonitor, &out.ServiceMonitor
		*out = new(MetricsServiceMonitorProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusRule != nil {
		in, out := &in.PrometheusRule, &out.PrometheusRule
		*out = new(MetricsPrometheusRuleProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsProperties.
func (in *MetricsProperties) DeepCopy() *MetricsProperties {
	if in == nil {
		return nil
	}
	out := new(MetricsProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServiceMonitorProperties) DeepCopyInto(out *MetricsServiceMonitorProperties) {
	*out = *in
	if in.Relabellings != nil {
		in, out := &in.Relabellings, &out.Relabellings
		*out = make([]monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetricRelabellings != nil {
		in, out := &in.MetricRelabellings, &out.MetricRelabellings
		*out = make([]monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodTargetLabels != nil {
		in, out := &in.PodTargetLabels, &out.PodTargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServiceMonitorProperties.
func (in *MetricsServiceMonitorProperties) DeepCopy() *MetricsServiceMonitorProperties {
	if in == nil {
		return nil
	}
	out := new(MetricsServiceMonitorProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceProperties) DeepCopyInto(out *PersistenceProperties) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceProperties.
func (in *PersistenceProperties) DeepCopy() *PersistenceProperties {
	if in == nil {
		return nil
	}
	out := new(PersistenceProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentinelProperties) DeepCopyInto(out *SentinelProperties) {
	*out = *in
	in.KubernetesContainerProperties.DeepCopyInto(&out.KubernetesContainerProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentinelProperties.
func (in *SentinelProperties) DeepCopy() *SentinelProperties {
	if in == nil {
		return nil
	}
	out := new(SentinelProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSProperties) DeepCopyInto(out *TLSProperties) {
	*out = *in
	if in.CertManager != nil {
		in, out := &in.CertManager, &out.CertManager
		*out = new(CertManagerProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSProperties.
func (in *TLSProperties) DeepCopy() *TLSProperties {
	if in == nil {
		return nil
	}
	out := new(TLSProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Valkey) DeepCopyInto(out *Valkey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Valkey.
func (in *Valkey) DeepCopy() *Valkey {
	if in == nil {
		return nil
	}
	out := new(Valkey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Valkey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValkeyList) DeepCopyInto(out *ValkeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Valkey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValkeyList.
func (in *ValkeyList) DeepCopy() *ValkeyList {
	if in == nil {
		return nil
	}
	out := new(ValkeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValkeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValkeySpec) DeepCopyInto(out *ValkeySpec) {
	*out = *in
	in.KubernetesPodProperties.DeepCopyInto(&out.KubernetesPodProperties)
	in.KubernetesContainerProperties.DeepCopyInto(&out.KubernetesContainerProperties)
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sentinel != nil {
		in, out := &in.Sentinel, &out.Sentinel
		*out = new(SentinelProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(BindingProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValkeySpec.
func (in *ValkeySpec) DeepCopy() *ValkeySpec {
	if in == nil {
		return nil
	}
	out := new(ValkeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValkeyStatus) DeepCopyInto(out *ValkeyStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValkeyStatus.
func (in *ValkeyStatus) DeepCopy() *ValkeyStatus {
	if in == nil {
		return nil
	}
	out := new(ValkeyStatus)
	in.DeepCopyInto(out)
	return out
}
