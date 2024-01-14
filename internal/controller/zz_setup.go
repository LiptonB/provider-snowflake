// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	databasegrant "github.com/LiptonB/provider-snowflake/internal/controller/grant/databasegrant"
	rolegrant "github.com/LiptonB/provider-snowflake/internal/controller/grant/rolegrant"
	schemagrant "github.com/LiptonB/provider-snowflake/internal/controller/grant/schemagrant"
	providerconfig "github.com/LiptonB/provider-snowflake/internal/controller/providerconfig"
	database "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/database"
	role "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/role"
	schema "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/schema"
	share "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/share"
	user "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/user"
	view "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/view"
	warehouse "github.com/LiptonB/provider-snowflake/internal/controller/snowflake/warehouse"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		databasegrant.Setup,
		rolegrant.Setup,
		schemagrant.Setup,
		providerconfig.Setup,
		database.Setup,
		role.Setup,
		schema.Setup,
		share.Setup,
		user.Setup,
		view.Setup,
		warehouse.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
