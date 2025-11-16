package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationAzureBlobCmd = &cobra.Command{
	Use:   "describe-location-azure-blob",
	Short: "Provides details about how an DataSync transfer location for Microsoft Azure Blob Storage is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationAzureBlobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeLocationAzureBlobCmd).Standalone()

		datasync_describeLocationAzureBlobCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of your Azure Blob Storage transfer location.")
		datasync_describeLocationAzureBlobCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_describeLocationAzureBlobCmd)
}
