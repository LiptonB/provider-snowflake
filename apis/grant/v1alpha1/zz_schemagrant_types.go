// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SchemaGrantInitParameters struct {

	// (Boolean) When this is set to true, multiple grants of the same type can be created.
	// When this is set to true, multiple grants of the same type can be created.
	EnableMultipleGrants *bool `json:"enableMultipleGrants,omitempty" tf:"enable_multiple_grants,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	// When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	OnAll *bool `json:"onAll,omitempty" tf:"on_all,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	// When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	OnFuture *bool `json:"onFuture,omitempty" tf:"on_future,omitempty"`

	// (String) The privilege to grant on the current or future schema. To grant all privileges, use the value ALL PRIVILEGES
	// The privilege to grant on the current or future schema. To grant all privileges, use the value `ALL PRIVILEGES`
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// (String) The name of the role to revert ownership to on destroy. Has no effect unless privilege is set to OWNERSHIP
	// The name of the role to revert ownership to on destroy. Has no effect unless `privilege` is set to `OWNERSHIP`
	RevertOwnershipToRoleName *string `json:"revertOwnershipToRoleName,omitempty" tf:"revert_ownership_to_role_name,omitempty"`

	// (Boolean) When this is set to true, allows the recipient role to grant the privileges to other roles.
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

type SchemaGrantObservation struct {

	// (String) The name of the database containing the schema on which to grant privileges.
	// The name of the database containing the schema on which to grant privileges.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// (Boolean) When this is set to true, multiple grants of the same type can be created.
	// When this is set to true, multiple grants of the same type can be created.
	EnableMultipleGrants *bool `json:"enableMultipleGrants,omitempty" tf:"enable_multiple_grants,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	// When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	OnAll *bool `json:"onAll,omitempty" tf:"on_all,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	// When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	OnFuture *bool `json:"onFuture,omitempty" tf:"on_future,omitempty"`

	// (String) The privilege to grant on the current or future schema. To grant all privileges, use the value ALL PRIVILEGES
	// The privilege to grant on the current or future schema. To grant all privileges, use the value `ALL PRIVILEGES`
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// (String) The name of the role to revert ownership to on destroy. Has no effect unless privilege is set to OWNERSHIP
	// The name of the role to revert ownership to on destroy. Has no effect unless `privilege` is set to `OWNERSHIP`
	RevertOwnershipToRoleName *string `json:"revertOwnershipToRoleName,omitempty" tf:"revert_ownership_to_role_name,omitempty"`

	// (Set of String) Grants privilege to these roles.
	// Grants privilege to these roles.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// (String) The name of the schema on which to grant privileges.
	// The name of the schema on which to grant privileges.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// (Set of String) Grants privilege to these shares (only valid if on_future and on_all are unset).
	// Grants privilege to these shares (only valid if on_future and on_all are unset).
	Shares []*string `json:"shares,omitempty" tf:"shares,omitempty"`

	// (Boolean) When this is set to true, allows the recipient role to grant the privileges to other roles.
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

type SchemaGrantParameters struct {

	// (String) The name of the database containing the schema on which to grant privileges.
	// The name of the database containing the schema on which to grant privileges.
	// +crossplane:generate:reference:type=github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Database
	// +crossplane:generate:reference:refFieldName=DatabaseRef
	// +crossplane:generate:reference:selectorFieldName=DatabaseSelector
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in snowflake to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in snowflake to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// (Boolean) When this is set to true, multiple grants of the same type can be created.
	// When this is set to true, multiple grants of the same type can be created.
	// +kubebuilder:validation:Optional
	EnableMultipleGrants *bool `json:"enableMultipleGrants,omitempty" tf:"enable_multiple_grants,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	// When this is set to true, apply this grant on all schemas in the given database. The schema_name and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	// +kubebuilder:validation:Optional
	OnAll *bool `json:"onAll,omitempty" tf:"on_all,omitempty"`

	// (Boolean) When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	// When this is set to true, apply this grant on all future schemas in the given database. The schema_name and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	// +kubebuilder:validation:Optional
	OnFuture *bool `json:"onFuture,omitempty" tf:"on_future,omitempty"`

	// (String) The privilege to grant on the current or future schema. To grant all privileges, use the value ALL PRIVILEGES
	// The privilege to grant on the current or future schema. To grant all privileges, use the value `ALL PRIVILEGES`
	// +kubebuilder:validation:Optional
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// (String) The name of the role to revert ownership to on destroy. Has no effect unless privilege is set to OWNERSHIP
	// The name of the role to revert ownership to on destroy. Has no effect unless `privilege` is set to `OWNERSHIP`
	// +kubebuilder:validation:Optional
	RevertOwnershipToRoleName *string `json:"revertOwnershipToRoleName,omitempty" tf:"revert_ownership_to_role_name,omitempty"`

	// References to Role in snowflake to populate roles.
	// +kubebuilder:validation:Optional
	RoleRefs []v1.Reference `json:"roleRefs,omitempty" tf:"-"`

	// Selector for a list of Role in snowflake to populate roles.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// (Set of String) Grants privilege to these roles.
	// Grants privilege to these roles.
	// +crossplane:generate:reference:type=github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Role
	// +crossplane:generate:reference:refFieldName=RoleRefs
	// +crossplane:generate:reference:selectorFieldName=RoleSelector
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// (String) The name of the schema on which to grant privileges.
	// The name of the schema on which to grant privileges.
	// +crossplane:generate:reference:type=github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Schema
	// +crossplane:generate:reference:refFieldName=SchemaRef
	// +crossplane:generate:reference:selectorFieldName=SchemaSelector
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Reference to a Schema in snowflake to populate schemaName.
	// +kubebuilder:validation:Optional
	SchemaRef *v1.Reference `json:"schemaRef,omitempty" tf:"-"`

	// Selector for a Schema in snowflake to populate schemaName.
	// +kubebuilder:validation:Optional
	SchemaSelector *v1.Selector `json:"schemaSelector,omitempty" tf:"-"`

	// References to Share in snowflake to populate shares.
	// +kubebuilder:validation:Optional
	ShareRefs []v1.Reference `json:"shareRefs,omitempty" tf:"-"`

	// Selector for a list of Share in snowflake to populate shares.
	// +kubebuilder:validation:Optional
	ShareSelector *v1.Selector `json:"shareSelector,omitempty" tf:"-"`

	// (Set of String) Grants privilege to these shares (only valid if on_future and on_all are unset).
	// Grants privilege to these shares (only valid if on_future and on_all are unset).
	// +crossplane:generate:reference:type=github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Share
	// +crossplane:generate:reference:refFieldName=ShareRefs
	// +crossplane:generate:reference:selectorFieldName=ShareSelector
	// +kubebuilder:validation:Optional
	Shares []*string `json:"shares,omitempty" tf:"shares,omitempty"`

	// (Boolean) When this is set to true, allows the recipient role to grant the privileges to other roles.
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	// +kubebuilder:validation:Optional
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

// SchemaGrantSpec defines the desired state of SchemaGrant
type SchemaGrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchemaGrantParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SchemaGrantInitParameters `json:"initProvider,omitempty"`
}

// SchemaGrantStatus defines the observed state of SchemaGrant.
type SchemaGrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchemaGrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaGrant is the Schema for the SchemaGrants API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,snowflake}
type SchemaGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemaGrantSpec   `json:"spec"`
	Status            SchemaGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaGrantList contains a list of SchemaGrants
type SchemaGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemaGrant `json:"items"`
}

// Repository type metadata.
var (
	SchemaGrant_Kind             = "SchemaGrant"
	SchemaGrant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SchemaGrant_Kind}.String()
	SchemaGrant_KindAPIVersion   = SchemaGrant_Kind + "." + CRDGroupVersion.String()
	SchemaGrant_GroupVersionKind = CRDGroupVersion.WithKind(SchemaGrant_Kind)
)

func init() {
	SchemeBuilder.Register(&SchemaGrant{}, &SchemaGrantList{})
}
