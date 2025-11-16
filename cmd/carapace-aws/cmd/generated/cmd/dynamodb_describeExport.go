package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeExportCmd = &cobra.Command{
	Use:   "describe-export",
	Short: "Describes an existing table export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_describeExportCmd).Standalone()

		dynamodb_describeExportCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) associated with the export.")
		dynamodb_describeExportCmd.MarkFlagRequired("export-arn")
	})
	dynamodbCmd.AddCommand(dynamodb_describeExportCmd)
}
