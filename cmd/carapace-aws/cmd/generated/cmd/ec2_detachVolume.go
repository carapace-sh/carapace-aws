package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachVolumeCmd = &cobra.Command{
	Use:   "detach-volume",
	Short: "Detaches an EBS volume from an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_detachVolumeCmd).Standalone()

		ec2_detachVolumeCmd.Flags().String("device", "", "The device name.")
		ec2_detachVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachVolumeCmd.Flags().Bool("force", false, "Forces detachment if the previous detachment attempt did not occur cleanly (for example, logging into an instance, unmounting the volume, and detaching normally).")
		ec2_detachVolumeCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_detachVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachVolumeCmd.Flags().Bool("no-force", false, "Forces detachment if the previous detachment attempt did not occur cleanly (for example, logging into an instance, unmounting the volume, and detaching normally).")
		ec2_detachVolumeCmd.Flags().String("volume-id", "", "The ID of the volume.")
		ec2_detachVolumeCmd.Flag("no-dry-run").Hidden = true
		ec2_detachVolumeCmd.Flag("no-force").Hidden = true
		ec2_detachVolumeCmd.MarkFlagRequired("volume-id")
	})
	ec2Cmd.AddCommand(ec2_detachVolumeCmd)
}
