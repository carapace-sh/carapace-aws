package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_startSnapshotCmd = &cobra.Command{
	Use:   "start-snapshot",
	Short: "Creates a new Amazon EBS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_startSnapshotCmd).Standalone()

	ebs_startSnapshotCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ebs_startSnapshotCmd.Flags().String("description", "", "A description for the snapshot.")
	ebs_startSnapshotCmd.Flags().Bool("encrypted", false, "Indicates whether to encrypt the snapshot.")
	ebs_startSnapshotCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the Key Management Service (KMS) key to be used to encrypt the snapshot.")
	ebs_startSnapshotCmd.Flags().Bool("no-encrypted", false, "Indicates whether to encrypt the snapshot.")
	ebs_startSnapshotCmd.Flags().String("parent-snapshot-id", "", "The ID of the parent snapshot.")
	ebs_startSnapshotCmd.Flags().String("tags", "", "The tags to apply to the snapshot.")
	ebs_startSnapshotCmd.Flags().String("timeout", "", "The amount of time (in minutes) after which the snapshot is automatically cancelled if:")
	ebs_startSnapshotCmd.Flags().String("volume-size", "", "The size of the volume, in GiB.")
	ebs_startSnapshotCmd.Flag("no-encrypted").Hidden = true
	ebs_startSnapshotCmd.MarkFlagRequired("volume-size")
	ebsCmd.AddCommand(ebs_startSnapshotCmd)
}
