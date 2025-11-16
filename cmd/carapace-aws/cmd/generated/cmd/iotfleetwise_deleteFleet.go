package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteFleetCmd = &cobra.Command{
	Use:   "delete-fleet",
	Short: "Deletes a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_deleteFleetCmd).Standalone()

		iotfleetwise_deleteFleetCmd.Flags().String("fleet-id", "", "The ID of the fleet to delete.")
		iotfleetwise_deleteFleetCmd.MarkFlagRequired("fleet-id")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteFleetCmd)
}
