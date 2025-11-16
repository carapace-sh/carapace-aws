package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_startDicomimportJobCmd = &cobra.Command{
	Use:   "start-dicomimport-job",
	Short: "Start importing bulk data into an `ACTIVE` data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_startDicomimportJobCmd).Standalone()

	medicalImaging_startDicomimportJobCmd.Flags().String("client-token", "", "A unique identifier for API idempotency.")
	medicalImaging_startDicomimportJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants permission to access medical imaging resources.")
	medicalImaging_startDicomimportJobCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_startDicomimportJobCmd.Flags().String("input-owner-account-id", "", "The account ID of the source S3 bucket owner.")
	medicalImaging_startDicomimportJobCmd.Flags().String("input-s3-uri", "", "The input prefix path for the S3 bucket that contains the DICOM files to be imported.")
	medicalImaging_startDicomimportJobCmd.Flags().String("job-name", "", "The import job name.")
	medicalImaging_startDicomimportJobCmd.Flags().String("output-s3-uri", "", "The output prefix of the S3 bucket to upload the results of the DICOM import job.")
	medicalImaging_startDicomimportJobCmd.MarkFlagRequired("client-token")
	medicalImaging_startDicomimportJobCmd.MarkFlagRequired("data-access-role-arn")
	medicalImaging_startDicomimportJobCmd.MarkFlagRequired("datastore-id")
	medicalImaging_startDicomimportJobCmd.MarkFlagRequired("input-s3-uri")
	medicalImaging_startDicomimportJobCmd.MarkFlagRequired("output-s3-uri")
	medicalImagingCmd.AddCommand(medicalImaging_startDicomimportJobCmd)
}
