package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_createCodeSigningConfigCmd = &cobra.Command{
	Use:   "create-code-signing-config",
	Short: "Creates a code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_createCodeSigningConfigCmd).Standalone()

	lambda_createCodeSigningConfigCmd.Flags().String("allowed-publishers", "", "Signing profiles for this code signing configuration.")
	lambda_createCodeSigningConfigCmd.Flags().String("code-signing-policies", "", "The code signing policies define the actions to take if the validation checks fail.")
	lambda_createCodeSigningConfigCmd.Flags().String("description", "", "Descriptive name for this code signing configuration.")
	lambda_createCodeSigningConfigCmd.Flags().String("tags", "", "A list of tags to add to the code signing configuration.")
	lambda_createCodeSigningConfigCmd.MarkFlagRequired("allowed-publishers")
	lambdaCmd.AddCommand(lambda_createCodeSigningConfigCmd)
}
