package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getHostReservationPurchasePreviewCmd = &cobra.Command{
	Use:   "get-host-reservation-purchase-preview",
	Short: "Preview a reservation purchase with configurations that match those of your Dedicated Host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getHostReservationPurchasePreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getHostReservationPurchasePreviewCmd).Standalone()

		ec2_getHostReservationPurchasePreviewCmd.Flags().String("host-id-set", "", "The IDs of the Dedicated Hosts with which the reservation is associated.")
		ec2_getHostReservationPurchasePreviewCmd.Flags().String("offering-id", "", "The offering ID of the reservation.")
		ec2_getHostReservationPurchasePreviewCmd.MarkFlagRequired("host-id-set")
		ec2_getHostReservationPurchasePreviewCmd.MarkFlagRequired("offering-id")
	})
	ec2Cmd.AddCommand(ec2_getHostReservationPurchasePreviewCmd)
}
