package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableSnapshotBlockPublicAccessCmd = &cobra.Command{
	Use:   "disable-snapshot-block-public-access",
	Short: "Disables the *block public access for snapshots* setting at the account level for the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableSnapshotBlockPublicAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableSnapshotBlockPublicAccessCmd).Standalone()

		ec2_disableSnapshotBlockPublicAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableSnapshotBlockPublicAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableSnapshotBlockPublicAccessCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disableSnapshotBlockPublicAccessCmd)
}
