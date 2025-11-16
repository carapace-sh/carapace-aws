package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateEntitlementCmd = &cobra.Command{
	Use:   "update-entitlement",
	Short: "Updates the specified entitlement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_updateEntitlementCmd).Standalone()

		appstream_updateEntitlementCmd.Flags().String("app-visibility", "", "Specifies whether all or only selected apps are entitled.")
		appstream_updateEntitlementCmd.Flags().String("attributes", "", "The attributes of the entitlement.")
		appstream_updateEntitlementCmd.Flags().String("description", "", "The description of the entitlement.")
		appstream_updateEntitlementCmd.Flags().String("name", "", "The name of the entitlement.")
		appstream_updateEntitlementCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
		appstream_updateEntitlementCmd.MarkFlagRequired("name")
		appstream_updateEntitlementCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_updateEntitlementCmd)
}
