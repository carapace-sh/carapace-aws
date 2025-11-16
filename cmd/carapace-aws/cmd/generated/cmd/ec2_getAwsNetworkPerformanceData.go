package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getAwsNetworkPerformanceDataCmd = &cobra.Command{
	Use:   "get-aws-network-performance-data",
	Short: "Gets network performance data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getAwsNetworkPerformanceDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getAwsNetworkPerformanceDataCmd).Standalone()

		ec2_getAwsNetworkPerformanceDataCmd.Flags().String("data-queries", "", "A list of network performance data queries.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().String("end-time", "", "The ending time for the performance data request.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getAwsNetworkPerformanceDataCmd.Flags().String("start-time", "", "The starting time for the performance data request.")
		ec2_getAwsNetworkPerformanceDataCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getAwsNetworkPerformanceDataCmd)
}
