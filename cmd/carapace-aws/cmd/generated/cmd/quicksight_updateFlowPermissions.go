package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateFlowPermissionsCmd = &cobra.Command{
	Use:   "update-flow-permissions",
	Short: "Updates permissions against principals on a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateFlowPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateFlowPermissionsCmd).Standalone()

		quicksight_updateFlowPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the flow you are updating permissions against.")
		quicksight_updateFlowPermissionsCmd.Flags().String("flow-id", "", "The unique identifier of the flow to update permissions for.")
		quicksight_updateFlowPermissionsCmd.Flags().String("grant-permissions", "", "The permissions that you want to grant on this flow.")
		quicksight_updateFlowPermissionsCmd.Flags().String("revoke-permissions", "", "The permissions that you want to revoke from this flow.")
		quicksight_updateFlowPermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateFlowPermissionsCmd.MarkFlagRequired("flow-id")
	})
	quicksightCmd.AddCommand(quicksight_updateFlowPermissionsCmd)
}
