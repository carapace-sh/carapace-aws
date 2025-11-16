package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getFunctionCmd = &cobra.Command{
	Use:   "get-function",
	Short: "Returns information about the function or function version, with a link to download the deployment package that's valid for 10 minutes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getFunctionCmd).Standalone()

	lambda_getFunctionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
	lambda_getFunctionCmd.Flags().String("qualifier", "", "Specify a version or alias to get details about a published version of the function.")
	lambda_getFunctionCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_getFunctionCmd)
}
