package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_exportTableToPointInTimeCmd = &cobra.Command{
	Use:   "export-table-to-point-in-time",
	Short: "Exports table data to an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_exportTableToPointInTimeCmd).Standalone()

	dynamodb_exportTableToPointInTimeCmd.Flags().String("client-token", "", "Providing a `ClientToken` makes the call to `ExportTableToPointInTimeInput` idempotent, meaning that multiple identical calls have the same effect as one single call.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("export-format", "", "The format for the exported data.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("export-time", "", "Time in the past from which to export table data, counted in seconds from the start of the Unix epoch.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("export-type", "", "Choice of whether to execute as a full export or incremental export.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("incremental-export-specification", "", "Optional object containing the parameters specific to an incremental export.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("s3-bucket", "", "The name of the Amazon S3 bucket to export the snapshot to.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("s3-bucket-owner", "", "The ID of the Amazon Web Services account that owns the bucket the export will be stored in.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("s3-prefix", "", "The Amazon S3 bucket prefix to use as the file name and path of the exported snapshot.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("s3-sse-algorithm", "", "Type of encryption used on the bucket where export data will be stored.")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("s3-sse-kms-key-id", "", "The ID of the KMS managed key used to encrypt the S3 bucket where export data will be stored (if applicable).")
	dynamodb_exportTableToPointInTimeCmd.Flags().String("table-arn", "", "The Amazon Resource Name (ARN) associated with the table to export.")
	dynamodb_exportTableToPointInTimeCmd.MarkFlagRequired("s3-bucket")
	dynamodb_exportTableToPointInTimeCmd.MarkFlagRequired("table-arn")
	dynamodbCmd.AddCommand(dynamodb_exportTableToPointInTimeCmd)
}
