package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachVolumeCmd = &cobra.Command{
	Use:   "attach-volume",
	Short: "Attaches an Amazon EBS volume to a `running` or `stopped` instance, and exposes it to the instance with the specified device name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_attachVolumeCmd).Standalone()

		ec2_attachVolumeCmd.Flags().String("device", "", "The device name (for example, `/dev/sdh` or `xvdh`).")
		ec2_attachVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachVolumeCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_attachVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachVolumeCmd.Flags().String("volume-id", "", "The ID of the EBS volume.")
		ec2_attachVolumeCmd.MarkFlagRequired("device")
		ec2_attachVolumeCmd.MarkFlagRequired("instance-id")
		ec2_attachVolumeCmd.Flag("no-dry-run").Hidden = true
		ec2_attachVolumeCmd.MarkFlagRequired("volume-id")
	})
	ec2Cmd.AddCommand(ec2_attachVolumeCmd)
}
