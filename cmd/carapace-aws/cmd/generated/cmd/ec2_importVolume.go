package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importVolumeCmd = &cobra.Command{
	Use:   "import-volume",
	Short: "This API action supports only single-volume VMs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_importVolumeCmd).Standalone()

		ec2_importVolumeCmd.Flags().String("availability-zone", "", "The Availability Zone for the resulting EBS volume.")
		ec2_importVolumeCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone for the resulting EBS volume.")
		ec2_importVolumeCmd.Flags().String("description", "", "A description of the volume.")
		ec2_importVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importVolumeCmd.Flags().String("image", "", "The disk image.")
		ec2_importVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importVolumeCmd.Flags().String("volume", "", "The volume size.")
		ec2_importVolumeCmd.MarkFlagRequired("image")
		ec2_importVolumeCmd.Flag("no-dry-run").Hidden = true
		ec2_importVolumeCmd.MarkFlagRequired("volume")
	})
	ec2Cmd.AddCommand(ec2_importVolumeCmd)
}
