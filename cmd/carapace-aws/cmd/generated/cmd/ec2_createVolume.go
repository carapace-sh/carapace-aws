package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVolumeCmd = &cobra.Command{
	Use:   "create-volume",
	Short: "Creates an EBS volume that can be attached to an instance in the same Availability Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVolumeCmd).Standalone()

	ec2_createVolumeCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone in which to create the volume.")
	ec2_createVolumeCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone in which to create the volume.")
	ec2_createVolumeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVolumeCmd.Flags().Bool("encrypted", false, "Indicates whether the volume should be encrypted.")
	ec2_createVolumeCmd.Flags().String("iops", "", "The number of I/O operations per second (IOPS) to provision for the volume.")
	ec2_createVolumeCmd.Flags().String("kms-key-id", "", "The identifier of the KMS key to use for Amazon EBS encryption.")
	ec2_createVolumeCmd.Flags().Bool("multi-attach-enabled", false, "Indicates whether to enable Amazon EBS Multi-Attach.")
	ec2_createVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVolumeCmd.Flags().Bool("no-encrypted", false, "Indicates whether the volume should be encrypted.")
	ec2_createVolumeCmd.Flags().Bool("no-multi-attach-enabled", false, "Indicates whether to enable Amazon EBS Multi-Attach.")
	ec2_createVolumeCmd.Flags().String("operator", "", "Reserved for internal use.")
	ec2_createVolumeCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost on which to create the volume.")
	ec2_createVolumeCmd.Flags().String("size", "", "The size of the volume, in GiBs.")
	ec2_createVolumeCmd.Flags().String("snapshot-id", "", "The snapshot from which to create the volume.")
	ec2_createVolumeCmd.Flags().String("tag-specifications", "", "The tags to apply to the volume during creation.")
	ec2_createVolumeCmd.Flags().String("throughput", "", "The throughput to provision for the volume, in MiB/s.")
	ec2_createVolumeCmd.Flags().String("volume-initialization-rate", "", "Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download the snapshot blocks from Amazon S3 to the volume.")
	ec2_createVolumeCmd.Flags().String("volume-type", "", "The volume type.")
	ec2_createVolumeCmd.Flag("no-dry-run").Hidden = true
	ec2_createVolumeCmd.Flag("no-encrypted").Hidden = true
	ec2_createVolumeCmd.Flag("no-multi-attach-enabled").Hidden = true
	ec2Cmd.AddCommand(ec2_createVolumeCmd)
}
