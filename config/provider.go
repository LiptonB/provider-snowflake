/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/LiptonB/provider-snowflake/config/database"
	"github.com/LiptonB/provider-snowflake/config/role"
	"github.com/LiptonB/provider-snowflake/config/schema"
	"github.com/LiptonB/provider-snowflake/config/share"
	"github.com/LiptonB/provider-snowflake/config/user"
	"github.com/LiptonB/provider-snowflake/config/view"
	"github.com/LiptonB/provider-snowflake/config/warehouse"
)

const (
	resourcePrefix = "snowflake"
	modulePath     = "github.com/LiptonB/provider-snowflake"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		role.Configure,
		schema.Configure,
		share.Configure,
		user.Configure,
		view.Configure,
		warehouse.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
