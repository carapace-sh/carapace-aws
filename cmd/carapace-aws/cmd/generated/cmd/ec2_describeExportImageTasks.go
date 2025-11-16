package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeExportImageTasksCmd = &cobra.Command{
	Use:   "describe-export-image-tasks",
	Short: "Describes the specified export image tasks or all of your export image tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeExportImageTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeExportImageTasksCmd).Standalone()

		ec2_describeExportImageTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeExportImageTasksCmd.Flags().String("export-image-task-ids", "", "The IDs of the export image tasks.")
		ec2_describeExportImageTasksCmd.Flags().String("filters", "", "Filter tasks using the `task-state` filter and one of the following values: `active`, `completed`, `deleting`, or `deleted`.")
		ec2_describeExportImageTasksCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describeExportImageTasksCmd.Flags().String("next-token", "", "A token that indicates the next page of results.")
		ec2_describeExportImageTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeExportImageTasksCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeExportImageTasksCmd)
}
