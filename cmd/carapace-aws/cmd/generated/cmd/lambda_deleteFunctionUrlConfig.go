package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteFunctionUrlConfigCmd = &cobra.Command{
	Use:   "delete-function-url-config",
	Short: "Deletes a Lambda function URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteFunctionUrlConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteFunctionUrlConfigCmd).Standalone()

		lambda_deleteFunctionUrlConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_deleteFunctionUrlConfigCmd.Flags().String("qualifier", "", "The alias name.")
		lambda_deleteFunctionUrlConfigCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_deleteFunctionUrlConfigCmd)
}
