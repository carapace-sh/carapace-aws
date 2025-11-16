package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteSnapshotCmd).Standalone()

	ec2_deleteSnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the EBS snapshot.")
	ec2_deleteSnapshotCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteSnapshotCmd.MarkFlagRequired("snapshot-id")
	ec2Cmd.AddCommand(ec2_deleteSnapshotCmd)
}
