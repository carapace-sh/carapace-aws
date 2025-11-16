package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationNfsCmd = &cobra.Command{
	Use:   "create-location-nfs",
	Short: "Creates a transfer *location* for a Network File System (NFS) file server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationNfsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_createLocationNfsCmd).Standalone()

		datasync_createLocationNfsCmd.Flags().String("mount-options", "", "Specifies the options that DataSync can use to mount your NFS file server.")
		datasync_createLocationNfsCmd.Flags().String("on-prem-config", "", "Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect to your NFS file server.")
		datasync_createLocationNfsCmd.Flags().String("server-hostname", "", "Specifies the DNS name or IP address (IPv4 or IPv6) of the NFS file server that your DataSync agent connects to.")
		datasync_createLocationNfsCmd.Flags().String("subdirectory", "", "Specifies the export path in your NFS file server that you want DataSync to mount.")
		datasync_createLocationNfsCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
		datasync_createLocationNfsCmd.MarkFlagRequired("on-prem-config")
		datasync_createLocationNfsCmd.MarkFlagRequired("server-hostname")
		datasync_createLocationNfsCmd.MarkFlagRequired("subdirectory")
	})
	datasyncCmd.AddCommand(datasync_createLocationNfsCmd)
}
