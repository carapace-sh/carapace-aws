package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteFleetCmd = &cobra.Command{
	Use:   "delete-fleet",
	Short: "Deletes a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteFleetCmd).Standalone()

	deadline_deleteFleetCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_deleteFleetCmd.Flags().String("farm-id", "", "The farm ID of the farm to remove from the fleet.")
	deadline_deleteFleetCmd.Flags().String("fleet-id", "", "The fleet ID of the fleet to delete.")
	deadline_deleteFleetCmd.MarkFlagRequired("farm-id")
	deadline_deleteFleetCmd.MarkFlagRequired("fleet-id")
	deadlineCmd.AddCommand(deadline_deleteFleetCmd)
}
