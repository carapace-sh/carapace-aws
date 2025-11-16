package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disassociateApplicationFromEntitlementCmd = &cobra.Command{
	Use:   "disassociate-application-from-entitlement",
	Short: "Deletes the specified application from the specified entitlement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disassociateApplicationFromEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_disassociateApplicationFromEntitlementCmd).Standalone()

		appstream_disassociateApplicationFromEntitlementCmd.Flags().String("application-identifier", "", "The identifier of the application to remove from the entitlement.")
		appstream_disassociateApplicationFromEntitlementCmd.Flags().String("entitlement-name", "", "The name of the entitlement.")
		appstream_disassociateApplicationFromEntitlementCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
		appstream_disassociateApplicationFromEntitlementCmd.MarkFlagRequired("application-identifier")
		appstream_disassociateApplicationFromEntitlementCmd.MarkFlagRequired("entitlement-name")
		appstream_disassociateApplicationFromEntitlementCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_disassociateApplicationFromEntitlementCmd)
}
