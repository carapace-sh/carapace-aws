package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listFunctionsByCodeSigningConfigCmd = &cobra.Command{
	Use:   "list-functions-by-code-signing-config",
	Short: "List the functions that use the specified code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listFunctionsByCodeSigningConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_listFunctionsByCodeSigningConfigCmd).Standalone()

		lambda_listFunctionsByCodeSigningConfigCmd.Flags().String("code-signing-config-arn", "", "The The Amazon Resource Name (ARN) of the code signing configuration.")
		lambda_listFunctionsByCodeSigningConfigCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
		lambda_listFunctionsByCodeSigningConfigCmd.Flags().String("max-items", "", "Maximum number of items to return.")
		lambda_listFunctionsByCodeSigningConfigCmd.MarkFlagRequired("code-signing-config-arn")
	})
	lambdaCmd.AddCommand(lambda_listFunctionsByCodeSigningConfigCmd)
}
