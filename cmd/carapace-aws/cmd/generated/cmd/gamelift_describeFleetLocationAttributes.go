package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetLocationAttributesCmd = &cobra.Command{
	Use:   "describe-fleet-location-attributes",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRetrieves information on a fleet's remote locations, including life-cycle status and any suspended fleet activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetLocationAttributesCmd).Standalone()

	gamelift_describeFleetLocationAttributesCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve remote locations for.")
	gamelift_describeFleetLocationAttributesCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describeFleetLocationAttributesCmd.Flags().String("locations", "", "A list of fleet locations to retrieve information for.")
	gamelift_describeFleetLocationAttributesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_describeFleetLocationAttributesCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeFleetLocationAttributesCmd)
}
