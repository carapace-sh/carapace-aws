package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVolumeCmd = &cobra.Command{
	Use:   "modify-volume",
	Short: "You can modify several parameters of an existing EBS volume, including volume size, volume type, and IOPS capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVolumeCmd).Standalone()

	ec2_modifyVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVolumeCmd.Flags().String("iops", "", "The target IOPS rate of the volume.")
	ec2_modifyVolumeCmd.Flags().Bool("multi-attach-enabled", false, "Specifies whether to enable Amazon EBS Multi-Attach.")
	ec2_modifyVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVolumeCmd.Flags().Bool("no-multi-attach-enabled", false, "Specifies whether to enable Amazon EBS Multi-Attach.")
	ec2_modifyVolumeCmd.Flags().String("size", "", "The target size of the volume, in GiB.")
	ec2_modifyVolumeCmd.Flags().String("throughput", "", "The target throughput of the volume, in MiB/s.")
	ec2_modifyVolumeCmd.Flags().String("volume-id", "", "The ID of the volume.")
	ec2_modifyVolumeCmd.Flags().String("volume-type", "", "The target EBS volume type of the volume.")
	ec2_modifyVolumeCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVolumeCmd.Flag("no-multi-attach-enabled").Hidden = true
	ec2_modifyVolumeCmd.MarkFlagRequired("volume-id")
	ec2Cmd.AddCommand(ec2_modifyVolumeCmd)
}
