package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetEventsCmd = &cobra.Command{
	Use:   "describe-fleet-events",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves entries from a fleet's event log.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetEventsCmd).Standalone()

	gamelift_describeFleetEventsCmd.Flags().String("end-time", "", "The most recent date to retrieve event logs for.")
	gamelift_describeFleetEventsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to get event logs for.")
	gamelift_describeFleetEventsCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describeFleetEventsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_describeFleetEventsCmd.Flags().String("start-time", "", "The earliest date to retrieve event logs for.")
	gamelift_describeFleetEventsCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeFleetEventsCmd)
}
