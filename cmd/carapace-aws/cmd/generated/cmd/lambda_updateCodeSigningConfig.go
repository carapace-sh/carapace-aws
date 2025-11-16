package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateCodeSigningConfigCmd = &cobra.Command{
	Use:   "update-code-signing-config",
	Short: "Update the code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateCodeSigningConfigCmd).Standalone()

	lambda_updateCodeSigningConfigCmd.Flags().String("allowed-publishers", "", "Signing profiles for this code signing configuration.")
	lambda_updateCodeSigningConfigCmd.Flags().String("code-signing-config-arn", "", "The The Amazon Resource Name (ARN) of the code signing configuration.")
	lambda_updateCodeSigningConfigCmd.Flags().String("code-signing-policies", "", "The code signing policy.")
	lambda_updateCodeSigningConfigCmd.Flags().String("description", "", "Descriptive name for this code signing configuration.")
	lambda_updateCodeSigningConfigCmd.MarkFlagRequired("code-signing-config-arn")
	lambdaCmd.AddCommand(lambda_updateCodeSigningConfigCmd)
}
