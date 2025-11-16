package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSnapshotAttributeCmd = &cobra.Command{
	Use:   "describe-snapshot-attribute",
	Short: "Describes the specified attribute of the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSnapshotAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSnapshotAttributeCmd).Standalone()

		ec2_describeSnapshotAttributeCmd.Flags().String("attribute", "", "The snapshot attribute you would like to view.")
		ec2_describeSnapshotAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSnapshotAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSnapshotAttributeCmd.Flags().String("snapshot-id", "", "The ID of the EBS snapshot.")
		ec2_describeSnapshotAttributeCmd.MarkFlagRequired("attribute")
		ec2_describeSnapshotAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_describeSnapshotAttributeCmd.MarkFlagRequired("snapshot-id")
	})
	ec2Cmd.AddCommand(ec2_describeSnapshotAttributeCmd)
}
