package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listVersionsByFunctionCmd = &cobra.Command{
	Use:   "list-versions-by-function",
	Short: "Returns a list of [versions](https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html), with the version-specific configuration of each.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listVersionsByFunctionCmd).Standalone()

	lambda_listVersionsByFunctionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_listVersionsByFunctionCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
	lambda_listVersionsByFunctionCmd.Flags().String("max-items", "", "The maximum number of versions to return.")
	lambda_listVersionsByFunctionCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_listVersionsByFunctionCmd)
}
