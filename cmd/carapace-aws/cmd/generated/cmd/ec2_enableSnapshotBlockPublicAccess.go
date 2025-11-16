package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableSnapshotBlockPublicAccessCmd = &cobra.Command{
	Use:   "enable-snapshot-block-public-access",
	Short: "Enables or modifies the *block public access for snapshots* setting at the account level for the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableSnapshotBlockPublicAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableSnapshotBlockPublicAccessCmd).Standalone()

		ec2_enableSnapshotBlockPublicAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableSnapshotBlockPublicAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableSnapshotBlockPublicAccessCmd.Flags().String("state", "", "The mode in which to enable block public access for snapshots for the Region.")
		ec2_enableSnapshotBlockPublicAccessCmd.Flag("no-dry-run").Hidden = true
		ec2_enableSnapshotBlockPublicAccessCmd.MarkFlagRequired("state")
	})
	ec2Cmd.AddCommand(ec2_enableSnapshotBlockPublicAccessCmd)
}
