package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeMacModificationTasksCmd = &cobra.Command{
	Use:   "describe-mac-modification-tasks",
	Short: "Describes a System Integrity Protection (SIP) modification task or volume ownership delegation task for an Amazon EC2 Mac instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeMacModificationTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeMacModificationTasksCmd).Standalone()

		ec2_describeMacModificationTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeMacModificationTasksCmd.Flags().String("filters", "", "Specifies one or more filters for the request:")
		ec2_describeMacModificationTasksCmd.Flags().String("mac-modification-task-ids", "", "The ID of task.")
		ec2_describeMacModificationTasksCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
		ec2_describeMacModificationTasksCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		ec2_describeMacModificationTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeMacModificationTasksCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeMacModificationTasksCmd)
}
