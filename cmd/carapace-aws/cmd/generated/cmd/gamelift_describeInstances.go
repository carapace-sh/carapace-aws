package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeInstancesCmd = &cobra.Command{
	Use:   "describe-instances",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves information about the EC2 instances in an Amazon GameLift Servers managed fleet, including instance ID, connection data, and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeInstancesCmd).Standalone()

		gamelift_describeInstancesCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve instance information for.")
		gamelift_describeInstancesCmd.Flags().String("instance-id", "", "A unique identifier for an instance to retrieve.")
		gamelift_describeInstancesCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeInstancesCmd.Flags().String("location", "", "The name of a location to retrieve instance information for, in the form of an Amazon Web Services Region code such as `us-west-2`.")
		gamelift_describeInstancesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_describeInstancesCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_describeInstancesCmd)
}
