package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionUrlConfigCmd = &cobra.Command{
	Use:   "get-function-url-config",
	Short: "Returns details about a Lambda function URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionUrlConfigCmd).Standalone()

	lambda_getFunctionUrlConfigCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_getFunctionUrlConfigCmd.Flags().String("qualifier", "", "The alias name.")
	lambda_getFunctionUrlConfigCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionUrlConfigCmd)
}
