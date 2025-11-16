package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_copyVolumesCmd = &cobra.Command{
	Use:   "copy-volumes",
	Short: "Creates a crash-consistent, point-in-time copy of an existing Amazon EBS volume within the same Availability Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_copyVolumesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_copyVolumesCmd).Standalone()

		ec2_copyVolumesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_copyVolumesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_copyVolumesCmd.Flags().String("iops", "", "The number of I/O operations per second (IOPS) to provision for the volume copy.")
		ec2_copyVolumesCmd.Flags().Bool("multi-attach-enabled", false, "Indicates whether to enable Amazon EBS Multi-Attach for the volume copy.")
		ec2_copyVolumesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_copyVolumesCmd.Flags().Bool("no-multi-attach-enabled", false, "Indicates whether to enable Amazon EBS Multi-Attach for the volume copy.")
		ec2_copyVolumesCmd.Flags().String("size", "", "The size of the volume copy, in GiBs.")
		ec2_copyVolumesCmd.Flags().String("source-volume-id", "", "The ID of the source EBS volume to copy.")
		ec2_copyVolumesCmd.Flags().String("tag-specifications", "", "The tags to apply to the volume copy during creation.")
		ec2_copyVolumesCmd.Flags().String("throughput", "", "The throughput to provision for the volume copy, in MiB/s.")
		ec2_copyVolumesCmd.Flags().String("volume-type", "", "The volume type for the volume copy.")
		ec2_copyVolumesCmd.Flag("no-dry-run").Hidden = true
		ec2_copyVolumesCmd.Flag("no-multi-attach-enabled").Hidden = true
		ec2_copyVolumesCmd.MarkFlagRequired("source-volume-id")
	})
	ec2Cmd.AddCommand(ec2_copyVolumesCmd)
}
