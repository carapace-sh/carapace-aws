package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeS3AccessPointAttachmentsCmd = &cobra.Command{
	Use:   "describe-s3-access-point-attachments",
	Short: "Describes one or more S3 access points attached to Amazon FSx volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeS3AccessPointAttachmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeS3AccessPointAttachmentsCmd).Standalone()

		fsx_describeS3AccessPointAttachmentsCmd.Flags().String("filters", "", "Enter a filter Name and Values pair to view a select set of S3 access point attachments.")
		fsx_describeS3AccessPointAttachmentsCmd.Flags().String("max-results", "", "")
		fsx_describeS3AccessPointAttachmentsCmd.Flags().String("names", "", "The names of the S3 access point attachments whose descriptions you want to retrieve.")
		fsx_describeS3AccessPointAttachmentsCmd.Flags().String("next-token", "", "")
	})
	fsxCmd.AddCommand(fsx_describeS3AccessPointAttachmentsCmd)
}
