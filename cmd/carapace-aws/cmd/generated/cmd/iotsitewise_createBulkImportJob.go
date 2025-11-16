package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createBulkImportJobCmd = &cobra.Command{
	Use:   "create-bulk-import-job",
	Short: "Defines a job to ingest data to IoT SiteWise from Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createBulkImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_createBulkImportJobCmd).Standalone()

		iotsitewise_createBulkImportJobCmd.Flags().String("adaptive-ingestion", "", "If set to true, ingest new data into IoT SiteWise storage.")
		iotsitewise_createBulkImportJobCmd.Flags().String("delete-files-after-import", "", "If set to true, your data files is deleted from S3, after ingestion into IoT SiteWise storage.")
		iotsitewise_createBulkImportJobCmd.Flags().String("error-report-location", "", "The Amazon S3 destination where errors associated with the job creation request are saved.")
		iotsitewise_createBulkImportJobCmd.Flags().String("files", "", "The files in the specified Amazon S3 bucket that contain your data.")
		iotsitewise_createBulkImportJobCmd.Flags().String("job-configuration", "", "Contains the configuration information of a job, such as the file format used to save data in Amazon S3.")
		iotsitewise_createBulkImportJobCmd.Flags().String("job-name", "", "The unique name that helps identify the job request.")
		iotsitewise_createBulkImportJobCmd.Flags().String("job-role-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the IAM role that allows IoT SiteWise to read Amazon S3 data.")
		iotsitewise_createBulkImportJobCmd.MarkFlagRequired("error-report-location")
		iotsitewise_createBulkImportJobCmd.MarkFlagRequired("files")
		iotsitewise_createBulkImportJobCmd.MarkFlagRequired("job-configuration")
		iotsitewise_createBulkImportJobCmd.MarkFlagRequired("job-name")
		iotsitewise_createBulkImportJobCmd.MarkFlagRequired("job-role-arn")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_createBulkImportJobCmd)
}
