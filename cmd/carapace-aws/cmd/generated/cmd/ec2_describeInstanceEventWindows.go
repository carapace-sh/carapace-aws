package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceEventWindowsCmd = &cobra.Command{
	Use:   "describe-instance-event-windows",
	Short: "Describes the specified event windows or all event windows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceEventWindowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeInstanceEventWindowsCmd).Standalone()

		ec2_describeInstanceEventWindowsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeInstanceEventWindowsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeInstanceEventWindowsCmd.Flags().String("instance-event-window-ids", "", "The IDs of the event windows.")
		ec2_describeInstanceEventWindowsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describeInstanceEventWindowsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		ec2_describeInstanceEventWindowsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeInstanceEventWindowsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeInstanceEventWindowsCmd)
}
