package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: "Lists completed exports within the past 90 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listExportsCmd).Standalone()

	dynamodb_listExportsCmd.Flags().String("max-results", "", "Maximum number of results to return per page.")
	dynamodb_listExportsCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to `ListExports`.")
	dynamodb_listExportsCmd.Flags().String("table-arn", "", "The Amazon Resource Name (ARN) associated with the exported table.")
	dynamodbCmd.AddCommand(dynamodb_listExportsCmd)
}
