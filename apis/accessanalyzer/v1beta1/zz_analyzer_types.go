/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AnalyzerInitParameters struct {

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of Analyzer. Valid values are ACCOUNT or ORGANIZATION. Defaults to ACCOUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AnalyzerObservation struct {

	// ARN of the Analyzer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Analyzer name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Type of Analyzer. Valid values are ACCOUNT or ORGANIZATION. Defaults to ACCOUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AnalyzerParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of Analyzer. Valid values are ACCOUNT or ORGANIZATION. Defaults to ACCOUNT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AnalyzerSpec defines the desired state of Analyzer
type AnalyzerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyzerParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AnalyzerInitParameters `json:"initProvider,omitempty"`
}

// AnalyzerStatus defines the observed state of Analyzer.
type AnalyzerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyzerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Analyzer is the Schema for the Analyzers API. Manages an Access Analyzer Analyzer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Analyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyzerSpec   `json:"spec"`
	Status            AnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyzerList contains a list of Analyzers
type AnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Analyzer `json:"items"`
}

// Repository type metadata.
var (
	Analyzer_Kind             = "Analyzer"
	Analyzer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Analyzer_Kind}.String()
	Analyzer_KindAPIVersion   = Analyzer_Kind + "." + CRDGroupVersion.String()
	Analyzer_GroupVersionKind = CRDGroupVersion.WithKind(Analyzer_Kind)
)

func init() {
	SchemeBuilder.Register(&Analyzer{}, &AnalyzerList{})
}
