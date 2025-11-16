package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFastSnapshotRestoresCmd = &cobra.Command{
	Use:   "describe-fast-snapshot-restores",
	Short: "Describes the state of fast snapshot restores for your snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFastSnapshotRestoresCmd).Standalone()

	ec2_describeFastSnapshotRestoresCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFastSnapshotRestoresCmd.Flags().String("filters", "", "The filters.")
	ec2_describeFastSnapshotRestoresCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeFastSnapshotRestoresCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeFastSnapshotRestoresCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFastSnapshotRestoresCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeFastSnapshotRestoresCmd)
}
