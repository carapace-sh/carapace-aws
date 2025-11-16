package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_getFlowPermissionsCmd = &cobra.Command{
	Use:   "get-flow-permissions",
	Short: "Get permissions for a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_getFlowPermissionsCmd).Standalone()

	quicksight_getFlowPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the flow that you are getting permissions for.")
	quicksight_getFlowPermissionsCmd.Flags().String("flow-id", "", "The unique identifier of the flow to get permissions from.")
	quicksight_getFlowPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_getFlowPermissionsCmd.MarkFlagRequired("flow-id")
	quicksightCmd.AddCommand(quicksight_getFlowPermissionsCmd)
}
