package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listImportsCmd = &cobra.Command{
	Use:   "list-imports",
	Short: "Lists completed imports within the past 90 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listImportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_listImportsCmd).Standalone()

		dynamodb_listImportsCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to `ListImports`.")
		dynamodb_listImportsCmd.Flags().String("page-size", "", "The number of `ImportSummary`objects returned in a single page.")
		dynamodb_listImportsCmd.Flags().String("table-arn", "", "The Amazon Resource Name (ARN) associated with the table that was imported to.")
	})
	dynamodbCmd.AddCommand(dynamodb_listImportsCmd)
}
