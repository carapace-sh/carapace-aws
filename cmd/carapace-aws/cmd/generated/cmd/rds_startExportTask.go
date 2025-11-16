package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_startExportTaskCmd = &cobra.Command{
	Use:   "start-export-task",
	Short: "Starts an export of DB snapshot or DB cluster data to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_startExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_startExportTaskCmd).Standalone()

		rds_startExportTaskCmd.Flags().String("export-only", "", "The data to be exported from the snapshot or cluster.")
		rds_startExportTaskCmd.Flags().String("export-task-identifier", "", "A unique identifier for the export task.")
		rds_startExportTaskCmd.Flags().String("iam-role-arn", "", "The name of the IAM role to use for writing to the Amazon S3 bucket when exporting a snapshot or cluster.")
		rds_startExportTaskCmd.Flags().String("kms-key-id", "", "The ID of the Amazon Web Services KMS key to use to encrypt the data exported to Amazon S3.")
		rds_startExportTaskCmd.Flags().String("s3-bucket-name", "", "The name of the Amazon S3 bucket to export the snapshot or cluster data to.")
		rds_startExportTaskCmd.Flags().String("s3-prefix", "", "The Amazon S3 bucket prefix to use as the file name and path of the exported data.")
		rds_startExportTaskCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) of the snapshot or cluster to export to Amazon S3.")
		rds_startExportTaskCmd.MarkFlagRequired("export-task-identifier")
		rds_startExportTaskCmd.MarkFlagRequired("iam-role-arn")
		rds_startExportTaskCmd.MarkFlagRequired("kms-key-id")
		rds_startExportTaskCmd.MarkFlagRequired("s3-bucket-name")
		rds_startExportTaskCmd.MarkFlagRequired("source-arn")
	})
	rdsCmd.AddCommand(rds_startExportTaskCmd)
}
