package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeGlobalTableCmd = &cobra.Command{
	Use:   "describe-global-table",
	Short: "Returns information about the specified global table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeGlobalTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_describeGlobalTableCmd).Standalone()

		dynamodb_describeGlobalTableCmd.Flags().String("global-table-name", "", "The name of the global table.")
		dynamodb_describeGlobalTableCmd.MarkFlagRequired("global-table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_describeGlobalTableCmd)
}
