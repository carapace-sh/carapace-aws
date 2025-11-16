package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "Returns a list of [aliases](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html) for a Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listAliasesCmd).Standalone()

	lambda_listAliasesCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_listAliasesCmd.Flags().String("function-version", "", "Specify a function version to only list aliases that invoke that version.")
	lambda_listAliasesCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
	lambda_listAliasesCmd.Flags().String("max-items", "", "Limit the number of aliases returned.")
	lambda_listAliasesCmd.MarkFlagRequired("function-name")
	lambdaCmd.AddCommand(lambda_listAliasesCmd)
}
