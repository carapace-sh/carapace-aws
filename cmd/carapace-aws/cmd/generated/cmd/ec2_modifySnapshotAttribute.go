package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifySnapshotAttributeCmd = &cobra.Command{
	Use:   "modify-snapshot-attribute",
	Short: "Adds or removes permission settings for the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifySnapshotAttributeCmd).Standalone()

	ec2_modifySnapshotAttributeCmd.Flags().String("attribute", "", "The snapshot attribute to modify.")
	ec2_modifySnapshotAttributeCmd.Flags().String("create-volume-permission", "", "A JSON representation of the snapshot attribute modification.")
	ec2_modifySnapshotAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySnapshotAttributeCmd.Flags().String("group-names", "", "The group to modify for the snapshot.")
	ec2_modifySnapshotAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifySnapshotAttributeCmd.Flags().String("operation-type", "", "The type of operation to perform to the attribute.")
	ec2_modifySnapshotAttributeCmd.Flags().String("snapshot-id", "", "The ID of the snapshot.")
	ec2_modifySnapshotAttributeCmd.Flags().String("user-ids", "", "The account ID to modify for the snapshot.")
	ec2_modifySnapshotAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2_modifySnapshotAttributeCmd.MarkFlagRequired("snapshot-id")
	ec2Cmd.AddCommand(ec2_modifySnapshotAttributeCmd)
}
