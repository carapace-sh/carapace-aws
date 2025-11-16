package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateActionTargetCmd = &cobra.Command{
	Use:   "update-action-target",
	Short: "Updates the name and description of a custom action target in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateActionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateActionTargetCmd).Standalone()

		securityhub_updateActionTargetCmd.Flags().String("action-target-arn", "", "The ARN of the custom action target to update.")
		securityhub_updateActionTargetCmd.Flags().String("description", "", "The updated description for the custom action target.")
		securityhub_updateActionTargetCmd.Flags().String("name", "", "The updated name of the custom action target.")
		securityhub_updateActionTargetCmd.MarkFlagRequired("action-target-arn")
	})
	securityhubCmd.AddCommand(securityhub_updateActionTargetCmd)
}
