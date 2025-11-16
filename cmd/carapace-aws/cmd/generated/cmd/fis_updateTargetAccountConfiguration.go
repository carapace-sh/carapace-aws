package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_updateTargetAccountConfigurationCmd = &cobra.Command{
	Use:   "update-target-account-configuration",
	Short: "Updates the target account configuration for the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_updateTargetAccountConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_updateTargetAccountConfigurationCmd).Standalone()

		fis_updateTargetAccountConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the target account.")
		fis_updateTargetAccountConfigurationCmd.Flags().String("description", "", "The description of the target account.")
		fis_updateTargetAccountConfigurationCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
		fis_updateTargetAccountConfigurationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role for the target account.")
		fis_updateTargetAccountConfigurationCmd.MarkFlagRequired("account-id")
		fis_updateTargetAccountConfigurationCmd.MarkFlagRequired("experiment-template-id")
	})
	fisCmd.AddCommand(fis_updateTargetAccountConfigurationCmd)
}
