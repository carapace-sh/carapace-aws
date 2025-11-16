package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifySnapshotTierCmd = &cobra.Command{
	Use:   "modify-snapshot-tier",
	Short: "Archives an Amazon EBS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifySnapshotTierCmd).Standalone()

	ec2_modifySnapshotTierCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySnapshotTierCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySnapshotTierCmd.Flags().String("snapshot-id", "", "The ID of the snapshot.")
	ec2_modifySnapshotTierCmd.Flags().String("storage-tier", "", "The name of the storage tier.")
	ec2_modifySnapshotTierCmd.Flag("no-dry-run").Hidden = true
	ec2_modifySnapshotTierCmd.MarkFlagRequired("snapshot-id")
	ec2Cmd.AddCommand(ec2_modifySnapshotTierCmd)
}
