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

type DelegatedAdministratorInitParameters struct {

	// The service principal of the AWS service for which you want to make the member account a delegated administrator.
	ServicePrincipal *string `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`
}

type DelegatedAdministratorObservation struct {

	// The account ID number of the member account in the organization to register as a delegated administrator.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The Amazon Resource Name (ARN) of the delegated administrator's account.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date when the account was made a delegated administrator.
	DelegationEnabledDate *string `json:"delegationEnabledDate,omitempty" tf:"delegation_enabled_date,omitempty"`

	// The email address that is associated with the delegated administrator's AWS account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The unique identifier (ID) of the delegated administrator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The method by which the delegated administrator's account joined the organization.
	JoinedMethod *string `json:"joinedMethod,omitempty" tf:"joined_method,omitempty"`

	// The date when the delegated administrator's account became a part of the organization.
	JoinedTimestamp *string `json:"joinedTimestamp,omitempty" tf:"joined_timestamp,omitempty"`

	// The friendly name of the delegated administrator's account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The service principal of the AWS service for which you want to make the member account a delegated administrator.
	ServicePrincipal *string `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`

	// The status of the delegated administrator's account in the organization.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DelegatedAdministratorParameters struct {

	// The account ID number of the member account in the organization to register as a delegated administrator.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/organizations/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in organizations to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in organizations to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The service principal of the AWS service for which you want to make the member account a delegated administrator.
	// +kubebuilder:validation:Optional
	ServicePrincipal *string `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`
}

// DelegatedAdministratorSpec defines the desired state of DelegatedAdministrator
type DelegatedAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DelegatedAdministratorParameters `json:"forProvider"`
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
	InitProvider DelegatedAdministratorInitParameters `json:"initProvider,omitempty"`
}

// DelegatedAdministratorStatus defines the observed state of DelegatedAdministrator.
type DelegatedAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DelegatedAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DelegatedAdministrator is the Schema for the DelegatedAdministrators API. Provides a resource to manage an AWS Organizations Delegated Administrator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DelegatedAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.servicePrincipal) || has(self.initProvider.servicePrincipal)",message="servicePrincipal is a required parameter"
	Spec   DelegatedAdministratorSpec   `json:"spec"`
	Status DelegatedAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DelegatedAdministratorList contains a list of DelegatedAdministrators
type DelegatedAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DelegatedAdministrator `json:"items"`
}

// Repository type metadata.
var (
	DelegatedAdministrator_Kind             = "DelegatedAdministrator"
	DelegatedAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DelegatedAdministrator_Kind}.String()
	DelegatedAdministrator_KindAPIVersion   = DelegatedAdministrator_Kind + "." + CRDGroupVersion.String()
	DelegatedAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(DelegatedAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&DelegatedAdministrator{}, &DelegatedAdministratorList{})
}
