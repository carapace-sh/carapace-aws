package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createEntitlementCmd = &cobra.Command{
	Use:   "create-entitlement",
	Short: "Creates a new entitlement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createEntitlementCmd).Standalone()

		appstream_createEntitlementCmd.Flags().String("app-visibility", "", "Specifies whether all or selected apps are entitled.")
		appstream_createEntitlementCmd.Flags().String("attributes", "", "The attributes of the entitlement.")
		appstream_createEntitlementCmd.Flags().String("description", "", "The description of the entitlement.")
		appstream_createEntitlementCmd.Flags().String("name", "", "The name of the entitlement.")
		appstream_createEntitlementCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
		appstream_createEntitlementCmd.MarkFlagRequired("app-visibility")
		appstream_createEntitlementCmd.MarkFlagRequired("attributes")
		appstream_createEntitlementCmd.MarkFlagRequired("name")
		appstream_createEntitlementCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_createEntitlementCmd)
}
