package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteCodeSigningConfigCmd = &cobra.Command{
	Use:   "delete-code-signing-config",
	Short: "Deletes the code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteCodeSigningConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteCodeSigningConfigCmd).Standalone()

		lambda_deleteCodeSigningConfigCmd.Flags().String("code-signing-config-arn", "", "The The Amazon Resource Name (ARN) of the code signing configuration.")
		lambda_deleteCodeSigningConfigCmd.MarkFlagRequired("code-signing-config-arn")
	})
	lambdaCmd.AddCommand(lambda_deleteCodeSigningConfigCmd)
}
