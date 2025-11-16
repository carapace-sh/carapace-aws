package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateFleetCmd = &cobra.Command{
	Use:   "update-fleet",
	Short: "Updates the description of an existing fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_updateFleetCmd).Standalone()

		iotfleetwise_updateFleetCmd.Flags().String("description", "", "An updated description of the fleet.")
		iotfleetwise_updateFleetCmd.Flags().String("fleet-id", "", "The ID of the fleet to update.")
		iotfleetwise_updateFleetCmd.MarkFlagRequired("fleet-id")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateFleetCmd)
}
