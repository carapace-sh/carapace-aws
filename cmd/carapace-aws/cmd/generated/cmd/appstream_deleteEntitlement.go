package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteEntitlementCmd = &cobra.Command{
	Use:   "delete-entitlement",
	Short: "Deletes the specified entitlement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteEntitlementCmd).Standalone()

		appstream_deleteEntitlementCmd.Flags().String("name", "", "The name of the entitlement.")
		appstream_deleteEntitlementCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
		appstream_deleteEntitlementCmd.MarkFlagRequired("name")
		appstream_deleteEntitlementCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_deleteEntitlementCmd)
}
