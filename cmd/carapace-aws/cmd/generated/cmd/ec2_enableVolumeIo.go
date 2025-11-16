package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableVolumeIoCmd = &cobra.Command{
	Use:   "enable-volume-io",
	Short: "Enables I/O operations for a volume that had I/O operations disabled because the data on the volume was potentially inconsistent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableVolumeIoCmd).Standalone()

	ec2_enableVolumeIoCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableVolumeIoCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableVolumeIoCmd.Flags().String("volume-id", "", "The ID of the volume.")
	ec2_enableVolumeIoCmd.Flag("no-dry-run").Hidden = true
	ec2_enableVolumeIoCmd.MarkFlagRequired("volume-id")
	ec2Cmd.AddCommand(ec2_enableVolumeIoCmd)
}
