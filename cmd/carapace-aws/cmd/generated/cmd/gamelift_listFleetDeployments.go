package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listFleetDeploymentsCmd = &cobra.Command{
	Use:   "list-fleet-deployments",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves a collection of container fleet deployments in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listFleetDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listFleetDeploymentsCmd).Standalone()

		gamelift_listFleetDeploymentsCmd.Flags().String("fleet-id", "", "A unique identifier for the container fleet.")
		gamelift_listFleetDeploymentsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listFleetDeploymentsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_listFleetDeploymentsCmd)
}
