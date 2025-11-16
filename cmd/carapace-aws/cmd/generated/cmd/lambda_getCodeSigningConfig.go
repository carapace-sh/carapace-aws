package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getCodeSigningConfigCmd = &cobra.Command{
	Use:   "get-code-signing-config",
	Short: "Returns information about the specified code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getCodeSigningConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_getCodeSigningConfigCmd).Standalone()

		lambda_getCodeSigningConfigCmd.Flags().String("code-signing-config-arn", "", "The The Amazon Resource Name (ARN) of the code signing configuration.")
		lambda_getCodeSigningConfigCmd.MarkFlagRequired("code-signing-config-arn")
	})
	lambdaCmd.AddCommand(lambda_getCodeSigningConfigCmd)
}
