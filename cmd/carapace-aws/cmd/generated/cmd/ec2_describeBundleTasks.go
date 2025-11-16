package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeBundleTasksCmd = &cobra.Command{
	Use:   "describe-bundle-tasks",
	Short: "Describes the specified bundle tasks or all of your bundle tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeBundleTasksCmd).Standalone()

	ec2_describeBundleTasksCmd.Flags().String("bundle-ids", "", "The bundle task IDs.")
	ec2_describeBundleTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeBundleTasksCmd.Flags().String("filters", "", "The filters.")
	ec2_describeBundleTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeBundleTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeBundleTasksCmd)
}
