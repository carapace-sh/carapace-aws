package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getFleetCmd = &cobra.Command{
	Use:   "get-fleet",
	Short: "Retrieves information about a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getFleetCmd).Standalone()

		iotfleetwise_getFleetCmd.Flags().String("fleet-id", "", "The ID of the fleet to retrieve information about.")
		iotfleetwise_getFleetCmd.MarkFlagRequired("fleet-id")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getFleetCmd)
}
