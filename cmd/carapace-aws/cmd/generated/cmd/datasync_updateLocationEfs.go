package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationEfsCmd = &cobra.Command{
	Use:   "update-location-efs",
	Short: "Modifies the following configuration parameters of the Amazon EFS transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationEfsCmd).Standalone()

	datasync_updateLocationEfsCmd.Flags().String("access-point-arn", "", "Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses to mount your Amazon EFS file system.")
	datasync_updateLocationEfsCmd.Flags().String("file-system-access-role-arn", "", "Specifies an Identity and Access Management (IAM) role that allows DataSync to access your Amazon EFS file system.")
	datasync_updateLocationEfsCmd.Flags().String("in-transit-encryption", "", "Specifies whether you want DataSync to use Transport Layer Security (TLS) 1.2 encryption when it transfers data to or from your Amazon EFS file system.")
	datasync_updateLocationEfsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the Amazon EFS transfer location that you're updating.")
	datasync_updateLocationEfsCmd.Flags().String("subdirectory", "", "Specifies a mount path for your Amazon EFS file system.")
	datasync_updateLocationEfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationEfsCmd)
}
