package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteActionTargetCmd = &cobra.Command{
	Use:   "delete-action-target",
	Short: "Deletes a custom action target from Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteActionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_deleteActionTargetCmd).Standalone()

		securityhub_deleteActionTargetCmd.Flags().String("action-target-arn", "", "The Amazon Resource Name (ARN) of the custom action target to delete.")
		securityhub_deleteActionTargetCmd.MarkFlagRequired("action-target-arn")
	})
	securityhubCmd.AddCommand(securityhub_deleteActionTargetCmd)
}
