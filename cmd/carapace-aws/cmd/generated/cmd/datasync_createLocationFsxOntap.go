package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationFsxOntapCmd = &cobra.Command{
	Use:   "create-location-fsx-ontap",
	Short: "Creates a transfer *location* for an Amazon FSx for NetApp ONTAP file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationFsxOntapCmd).Standalone()

	datasync_createLocationFsxOntapCmd.Flags().String("protocol", "", "")
	datasync_createLocationFsxOntapCmd.Flags().String("security-group-arns", "", "Specifies the Amazon EC2 security groups that provide access to your file system's preferred subnet.")
	datasync_createLocationFsxOntapCmd.Flags().String("storage-virtual-machine-arn", "", "Specifies the ARN of the storage virtual machine (SVM) in your file system where you want to copy data to or from.")
	datasync_createLocationFsxOntapCmd.Flags().String("subdirectory", "", "Specifies a path to the file share in the SVM where you want to transfer data to or from.")
	datasync_createLocationFsxOntapCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createLocationFsxOntapCmd.MarkFlagRequired("protocol")
	datasync_createLocationFsxOntapCmd.MarkFlagRequired("security-group-arns")
	datasync_createLocationFsxOntapCmd.MarkFlagRequired("storage-virtual-machine-arn")
	datasyncCmd.AddCommand(datasync_createLocationFsxOntapCmd)
}
