package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getSnapshotBlockPublicAccessStateCmd = &cobra.Command{
	Use:   "get-snapshot-block-public-access-state",
	Short: "Gets the current state of *block public access for snapshots* setting for the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getSnapshotBlockPublicAccessStateCmd).Standalone()

	ec2_getSnapshotBlockPublicAccessStateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSnapshotBlockPublicAccessStateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSnapshotBlockPublicAccessStateCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getSnapshotBlockPublicAccessStateCmd)
}
