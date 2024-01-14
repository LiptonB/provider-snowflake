package grant

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("snowflake_database_grant", func(r *config.Resource) {
		r.ShortGroup = "grant"
		r.Kind = "DatabaseGrant"
		r.References = config.References{
			"database_name": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Database",
				RefFieldName:      "DatabaseRef",
				SelectorFieldName: "DatabaseSelector",
			},
			"roles": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Role",
				RefFieldName:      "RoleRefs",
				SelectorFieldName: "RoleSelector",
			},
			"shares": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Share",
				RefFieldName:      "ShareRefs",
				SelectorFieldName: "ShareSelector",
			},
		}
	})

	p.AddResourceConfigurator("snowflake_schema_grant", func(r *config.Resource) {
		r.ShortGroup = "grant"
		r.Kind = "SchemaGrant"
		r.References = config.References{
			"database_name": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Database",
				RefFieldName:      "DatabaseRef",
				SelectorFieldName: "DatabaseSelector",
			},
			"schema_name": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Schema",
				RefFieldName:      "SchemaRef",
				SelectorFieldName: "SchemaSelector",
			},
			"roles": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Role",
				RefFieldName:      "RoleRefs",
				SelectorFieldName: "RoleSelector",
			},
			"shares": {
				Type:              "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Share",
				RefFieldName:      "ShareRefs",
				SelectorFieldName: "ShareSelector",
			},
		}
	})
}
