package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getFleetCmd = &cobra.Command{
	Use:   "get-fleet",
	Short: "Get a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getFleetCmd).Standalone()

	deadline_getFleetCmd.Flags().String("farm-id", "", "The farm ID of the farm in the fleet.")
	deadline_getFleetCmd.Flags().String("fleet-id", "", "The fleet ID of the fleet to get.")
	deadline_getFleetCmd.MarkFlagRequired("farm-id")
	deadline_getFleetCmd.MarkFlagRequired("fleet-id")
	deadlineCmd.AddCommand(deadline_getFleetCmd)
}
