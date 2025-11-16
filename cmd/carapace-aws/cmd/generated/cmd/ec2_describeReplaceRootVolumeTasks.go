package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeReplaceRootVolumeTasksCmd = &cobra.Command{
	Use:   "describe-replace-root-volume-tasks",
	Short: "Describes a root volume replacement task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeReplaceRootVolumeTasksCmd).Standalone()

	ec2_describeReplaceRootVolumeTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReplaceRootVolumeTasksCmd.Flags().String("filters", "", "Filter to use:")
	ec2_describeReplaceRootVolumeTasksCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeReplaceRootVolumeTasksCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeReplaceRootVolumeTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReplaceRootVolumeTasksCmd.Flags().String("replace-root-volume-task-ids", "", "The ID of the root volume replacement task to view.")
	ec2_describeReplaceRootVolumeTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeReplaceRootVolumeTasksCmd)
}
