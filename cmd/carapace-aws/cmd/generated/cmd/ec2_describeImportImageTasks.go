package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImportImageTasksCmd = &cobra.Command{
	Use:   "describe-import-image-tasks",
	Short: "Displays details about an import virtual machine or import snapshot tasks that are already created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImportImageTasksCmd).Standalone()

	ec2_describeImportImageTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImportImageTasksCmd.Flags().String("filters", "", "Filter tasks using the `task-state` filter and one of the following values: `active`, `completed`, `deleting`, or `deleted`.")
	ec2_describeImportImageTasksCmd.Flags().String("import-task-ids", "", "The IDs of the import image tasks.")
	ec2_describeImportImageTasksCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeImportImageTasksCmd.Flags().String("next-token", "", "A token that indicates the next page of results.")
	ec2_describeImportImageTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImportImageTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImportImageTasksCmd)
}
