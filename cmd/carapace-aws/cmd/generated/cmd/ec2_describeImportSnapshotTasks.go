package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImportSnapshotTasksCmd = &cobra.Command{
	Use:   "describe-import-snapshot-tasks",
	Short: "Describes your import snapshot tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImportSnapshotTasksCmd).Standalone()

	ec2_describeImportSnapshotTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImportSnapshotTasksCmd.Flags().String("filters", "", "The filters.")
	ec2_describeImportSnapshotTasksCmd.Flags().String("import-task-ids", "", "A list of import snapshot task IDs.")
	ec2_describeImportSnapshotTasksCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeImportSnapshotTasksCmd.Flags().String("next-token", "", "A token that indicates the next page of results.")
	ec2_describeImportSnapshotTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImportSnapshotTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImportSnapshotTasksCmd)
}
