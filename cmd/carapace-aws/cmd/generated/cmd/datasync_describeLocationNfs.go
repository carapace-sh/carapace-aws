package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationNfsCmd = &cobra.Command{
	Use:   "describe-location-nfs",
	Short: "Provides details about how an DataSync transfer location for a Network File System (NFS) file server is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationNfsCmd).Standalone()

	datasync_describeLocationNfsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the NFS location that you want information about.")
	datasync_describeLocationNfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationNfsCmd)
}
