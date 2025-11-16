package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeImportCmd = &cobra.Command{
	Use:   "describe-import",
	Short: "Represents the properties of the import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeImportCmd).Standalone()

	dynamodb_describeImportCmd.Flags().String("import-arn", "", "The Amazon Resource Name (ARN) associated with the table you're importing to.")
	dynamodb_describeImportCmd.MarkFlagRequired("import-arn")
	dynamodbCmd.AddCommand(dynamodb_describeImportCmd)
}
