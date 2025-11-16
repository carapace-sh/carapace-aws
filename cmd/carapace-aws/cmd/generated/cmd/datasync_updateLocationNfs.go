package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationNfsCmd = &cobra.Command{
	Use:   "update-location-nfs",
	Short: "Modifies the following configuration parameters of the Network File System (NFS) transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationNfsCmd).Standalone()

	datasync_updateLocationNfsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the NFS transfer location that you want to update.")
	datasync_updateLocationNfsCmd.Flags().String("mount-options", "", "")
	datasync_updateLocationNfsCmd.Flags().String("on-prem-config", "", "")
	datasync_updateLocationNfsCmd.Flags().String("server-hostname", "", "Specifies the DNS name or IP address (IPv4 or IPv6) of the NFS file server that your DataSync agent connects to.")
	datasync_updateLocationNfsCmd.Flags().String("subdirectory", "", "Specifies the export path in your NFS file server that you want DataSync to mount.")
	datasync_updateLocationNfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationNfsCmd)
}
