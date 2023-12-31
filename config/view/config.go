package view

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("snowflake_view", func(r *config.Resource) {
		r.References = config.References{
			"database": {
				Type: "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Database",
			},
			"schema": {
				Type: "github.com/LiptonB/provider-snowflake/apis/snowflake/v1alpha1.Schema",
			},
		}
	})
}
