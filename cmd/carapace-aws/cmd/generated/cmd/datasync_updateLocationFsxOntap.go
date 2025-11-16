package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationFsxOntapCmd = &cobra.Command{
	Use:   "update-location-fsx-ontap",
	Short: "Modifies the following configuration parameters of the Amazon FSx for NetApp ONTAP transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationFsxOntapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_updateLocationFsxOntapCmd).Standalone()

		datasync_updateLocationFsxOntapCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for ONTAP transfer location that you're updating.")
		datasync_updateLocationFsxOntapCmd.Flags().String("protocol", "", "Specifies the data transfer protocol that DataSync uses to access your Amazon FSx file system.")
		datasync_updateLocationFsxOntapCmd.Flags().String("subdirectory", "", "Specifies a path to the file share in the storage virtual machine (SVM) where you want to transfer data to or from.")
		datasync_updateLocationFsxOntapCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_updateLocationFsxOntapCmd)
}
