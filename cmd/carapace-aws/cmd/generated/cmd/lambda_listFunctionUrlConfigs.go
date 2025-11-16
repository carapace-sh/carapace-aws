package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listFunctionUrlConfigsCmd = &cobra.Command{
	Use:   "list-function-url-configs",
	Short: "Returns a list of Lambda function URLs for the specified function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listFunctionUrlConfigsCmd).Standalone()

	lambda_listFunctionUrlConfigsCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_listFunctionUrlConfigsCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
	lambda_listFunctionUrlConfigsCmd.Flags().String("max-items", "", "The maximum number of function URLs to return in the response.")
	lambda_listFunctionUrlConfigsCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_listFunctionUrlConfigsCmd)
}
