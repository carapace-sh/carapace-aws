package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetDeploymentCmd = &cobra.Command{
	Use:   "describe-fleet-deployment",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves information about a managed container fleet deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetDeploymentCmd).Standalone()

	gamelift_describeFleetDeploymentCmd.Flags().String("deployment-id", "", "A unique identifier for the deployment to return information for.")
	gamelift_describeFleetDeploymentCmd.Flags().String("fleet-id", "", "A unique identifier for the container fleet.")
	gamelift_describeFleetDeploymentCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeFleetDeploymentCmd)
}
