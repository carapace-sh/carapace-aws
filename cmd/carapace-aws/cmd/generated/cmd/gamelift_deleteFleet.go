package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteFleetCmd = &cobra.Command{
	Use:   "delete-fleet",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nDeletes all resources and information related to a fleet and shuts down any currently running fleet instances, including those in remote locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteFleetCmd).Standalone()

	gamelift_deleteFleetCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to be deleted.")
	gamelift_deleteFleetCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_deleteFleetCmd)
}
