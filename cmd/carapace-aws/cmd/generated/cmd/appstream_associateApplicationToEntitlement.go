package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_associateApplicationToEntitlementCmd = &cobra.Command{
	Use:   "associate-application-to-entitlement",
	Short: "Associates an application to entitle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_associateApplicationToEntitlementCmd).Standalone()

	appstream_associateApplicationToEntitlementCmd.Flags().String("application-identifier", "", "The identifier of the application.")
	appstream_associateApplicationToEntitlementCmd.Flags().String("entitlement-name", "", "The name of the entitlement.")
	appstream_associateApplicationToEntitlementCmd.Flags().String("stack-name", "", "The name of the stack.")
	appstream_associateApplicationToEntitlementCmd.MarkFlagRequired("application-identifier")
	appstream_associateApplicationToEntitlementCmd.MarkFlagRequired("entitlement-name")
	appstream_associateApplicationToEntitlementCmd.MarkFlagRequired("stack-name")
	appstreamCmd.AddCommand(appstream_associateApplicationToEntitlementCmd)
}
