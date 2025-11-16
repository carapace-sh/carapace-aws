package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkInsightsPathCmd = &cobra.Command{
	Use:   "delete-network-insights-path",
	Short: "Deletes the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkInsightsPathCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteNetworkInsightsPathCmd).Standalone()

		ec2_deleteNetworkInsightsPathCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNetworkInsightsPathCmd.Flags().String("network-insights-path-id", "", "The ID of the path.")
		ec2_deleteNetworkInsightsPathCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteNetworkInsightsPathCmd.MarkFlagRequired("network-insights-path-id")
		ec2_deleteNetworkInsightsPathCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteNetworkInsightsPathCmd)
}
