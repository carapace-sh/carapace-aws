package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_createTargetAccountConfigurationCmd = &cobra.Command{
	Use:   "create-target-account-configuration",
	Short: "Creates a target account configuration for the experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_createTargetAccountConfigurationCmd).Standalone()

	fis_createTargetAccountConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the target account.")
	fis_createTargetAccountConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	fis_createTargetAccountConfigurationCmd.Flags().String("description", "", "The description of the target account.")
	fis_createTargetAccountConfigurationCmd.Flags().String("experiment-template-id", "", "The experiment template ID.")
	fis_createTargetAccountConfigurationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role for the target account.")
	fis_createTargetAccountConfigurationCmd.MarkFlagRequired("account-id")
	fis_createTargetAccountConfigurationCmd.MarkFlagRequired("experiment-template-id")
	fis_createTargetAccountConfigurationCmd.MarkFlagRequired("role-arn")
	fisCmd.AddCommand(fis_createTargetAccountConfigurationCmd)
}
