package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_renewOfferingCmd = &cobra.Command{
	Use:   "renew-offering",
	Short: "Explicitly sets the quantity of devices to renew for an offering, starting from the `effectiveDate` of the next period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_renewOfferingCmd).Standalone()

	devicefarm_renewOfferingCmd.Flags().String("offering-id", "", "The ID of a request to renew an offering.")
	devicefarm_renewOfferingCmd.Flags().String("quantity", "", "The quantity requested in an offering renewal.")
	devicefarm_renewOfferingCmd.MarkFlagRequired("offering-id")
	devicefarm_renewOfferingCmd.MarkFlagRequired("quantity")
	devicefarmCmd.AddCommand(devicefarm_renewOfferingCmd)
}
