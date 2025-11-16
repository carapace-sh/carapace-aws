package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_copySnapshotCmd = &cobra.Command{
	Use:   "copy-snapshot",
	Short: "Creates an exact copy of an Amazon EBS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_copySnapshotCmd).Standalone()

	ec2_copySnapshotCmd.Flags().String("completion-duration-minutes", "", "Not supported when copying snapshots to or from Local Zones or Outposts.")
	ec2_copySnapshotCmd.Flags().String("description", "", "A description for the EBS snapshot.")
	ec2_copySnapshotCmd.Flags().String("destination-availability-zone", "", "The Local Zone, for example, `cn-north-1-pkx-1a` to which to copy the snapshot.")
	ec2_copySnapshotCmd.Flags().String("destination-outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost to which to copy the snapshot.")
	ec2_copySnapshotCmd.Flags().String("destination-region", "", "The destination Region to use in the `PresignedUrl` parameter of a snapshot copy operation.")
	ec2_copySnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_copySnapshotCmd.Flags().Bool("encrypted", false, "To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.")
	ec2_copySnapshotCmd.Flags().String("kms-key-id", "", "The identifier of the KMS key to use for Amazon EBS encryption.")
	ec2_copySnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_copySnapshotCmd.Flags().Bool("no-encrypted", false, "To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.")
	ec2_copySnapshotCmd.Flags().String("presigned-url", "", "When you copy an encrypted source snapshot using the Amazon EC2 Query API, you must supply a pre-signed URL.")
	ec2_copySnapshotCmd.Flags().String("source-region", "", "The ID of the Region that contains the snapshot to be copied.")
	ec2_copySnapshotCmd.Flags().String("source-snapshot-id", "", "The ID of the EBS snapshot to copy.")
	ec2_copySnapshotCmd.Flags().String("tag-specifications", "", "The tags to apply to the new snapshot.")
	ec2_copySnapshotCmd.Flag("no-dry-run").Hidden = true
	ec2_copySnapshotCmd.Flag("no-encrypted").Hidden = true
	ec2_copySnapshotCmd.MarkFlagRequired("source-region")
	ec2_copySnapshotCmd.MarkFlagRequired("source-snapshot-id")
	ec2Cmd.AddCommand(ec2_copySnapshotCmd)
}
