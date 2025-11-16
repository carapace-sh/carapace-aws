package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listComputeCmd = &cobra.Command{
	Use:   "list-compute",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves information on the compute resources in an Amazon GameLift Servers fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listComputeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listComputeCmd).Standalone()

		gamelift_listComputeCmd.Flags().String("compute-status", "", "The status of computes in a managed container fleet, based on the success of the latest update deployment.")
		gamelift_listComputeCmd.Flags().String("container-group-definition-name", "", "For computes in a managed container fleet, the name of the deployed container group definition.")
		gamelift_listComputeCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve compute resources for.")
		gamelift_listComputeCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listComputeCmd.Flags().String("location", "", "The name of a location to retrieve compute resources for.")
		gamelift_listComputeCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_listComputeCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_listComputeCmd)
}
