package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_importTableCmd = &cobra.Command{
	Use:   "import-table",
	Short: "Imports table data from an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_importTableCmd).Standalone()

	dynamodb_importTableCmd.Flags().String("client-token", "", "Providing a `ClientToken` makes the call to `ImportTableInput` idempotent, meaning that multiple identical calls have the same effect as one single call.")
	dynamodb_importTableCmd.Flags().String("input-compression-type", "", "Type of compression to be used on the input coming from the imported table.")
	dynamodb_importTableCmd.Flags().String("input-format", "", "The format of the source data.")
	dynamodb_importTableCmd.Flags().String("input-format-options", "", "Additional properties that specify how the input is formatted,")
	dynamodb_importTableCmd.Flags().String("s3-bucket-source", "", "The S3 bucket that provides the source for the import.")
	dynamodb_importTableCmd.Flags().String("table-creation-parameters", "", "Parameters for the table to import the data into.")
	dynamodb_importTableCmd.MarkFlagRequired("input-format")
	dynamodb_importTableCmd.MarkFlagRequired("s3-bucket-source")
	dynamodb_importTableCmd.MarkFlagRequired("table-creation-parameters")
	dynamodbCmd.AddCommand(dynamodb_importTableCmd)
}
