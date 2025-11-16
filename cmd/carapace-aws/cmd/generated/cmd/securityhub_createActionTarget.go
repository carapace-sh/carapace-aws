package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createActionTargetCmd = &cobra.Command{
	Use:   "create-action-target",
	Short: "Creates a custom action target in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createActionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createActionTargetCmd).Standalone()

		securityhub_createActionTargetCmd.Flags().String("description", "", "The description for the custom action target.")
		securityhub_createActionTargetCmd.Flags().String("id", "", "The ID for the custom action target.")
		securityhub_createActionTargetCmd.Flags().String("name", "", "The name of the custom action target.")
		securityhub_createActionTargetCmd.MarkFlagRequired("description")
		securityhub_createActionTargetCmd.MarkFlagRequired("id")
		securityhub_createActionTargetCmd.MarkFlagRequired("name")
	})
	securityhubCmd.AddCommand(securityhub_createActionTargetCmd)
}
