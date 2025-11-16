package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkInsightsPathCmd = &cobra.Command{
	Use:   "create-network-insights-path",
	Short: "Creates a path to analyze for reachability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkInsightsPathCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createNetworkInsightsPathCmd).Standalone()

		ec2_createNetworkInsightsPathCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createNetworkInsightsPathCmd.Flags().String("destination", "", "The ID or ARN of the destination.")
		ec2_createNetworkInsightsPathCmd.Flags().String("destination-ip", "", "The IP address of the destination.")
		ec2_createNetworkInsightsPathCmd.Flags().String("destination-port", "", "The destination port.")
		ec2_createNetworkInsightsPathCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInsightsPathCmd.Flags().String("filter-at-destination", "", "Scopes the analysis to network paths that match specific filters at the destination.")
		ec2_createNetworkInsightsPathCmd.Flags().String("filter-at-source", "", "Scopes the analysis to network paths that match specific filters at the source.")
		ec2_createNetworkInsightsPathCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInsightsPathCmd.Flags().String("protocol", "", "The protocol.")
		ec2_createNetworkInsightsPathCmd.Flags().String("source", "", "The ID or ARN of the source.")
		ec2_createNetworkInsightsPathCmd.Flags().String("source-ip", "", "The IP address of the source.")
		ec2_createNetworkInsightsPathCmd.Flags().String("tag-specifications", "", "The tags to add to the path.")
		ec2_createNetworkInsightsPathCmd.MarkFlagRequired("client-token")
		ec2_createNetworkInsightsPathCmd.Flag("no-dry-run").Hidden = true
		ec2_createNetworkInsightsPathCmd.MarkFlagRequired("protocol")
		ec2_createNetworkInsightsPathCmd.MarkFlagRequired("source")
	})
	ec2Cmd.AddCommand(ec2_createNetworkInsightsPathCmd)
}
