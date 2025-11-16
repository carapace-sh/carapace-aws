package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_updateJobShipmentStateCmd = &cobra.Command{
	Use:   "update-job-shipment-state",
	Short: "Updates the state when a shipment state changes to a different state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_updateJobShipmentStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_updateJobShipmentStateCmd).Standalone()

		snowball_updateJobShipmentStateCmd.Flags().String("job-id", "", "The job ID of the job whose shipment date you want to update, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
		snowball_updateJobShipmentStateCmd.Flags().String("shipment-state", "", "The state of a device when it is being shipped.")
		snowball_updateJobShipmentStateCmd.MarkFlagRequired("job-id")
		snowball_updateJobShipmentStateCmd.MarkFlagRequired("shipment-state")
	})
	snowballCmd.AddCommand(snowball_updateJobShipmentStateCmd)
}
