package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_purchaseOfferingCmd = &cobra.Command{
	Use:   "purchase-offering",
	Short: "Submits a request to purchase an offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_purchaseOfferingCmd).Standalone()

	mediaconnect_purchaseOfferingCmd.Flags().String("offering-arn", "", "The Amazon Resource Name (ARN) of the offering.")
	mediaconnect_purchaseOfferingCmd.Flags().String("reservation-name", "", "The name that you want to use for the reservation.")
	mediaconnect_purchaseOfferingCmd.Flags().String("start", "", "The date and time that you want the reservation to begin, in Coordinated Universal Time (UTC).")
	mediaconnect_purchaseOfferingCmd.MarkFlagRequired("offering-arn")
	mediaconnect_purchaseOfferingCmd.MarkFlagRequired("reservation-name")
	mediaconnect_purchaseOfferingCmd.MarkFlagRequired("start")
	mediaconnectCmd.AddCommand(mediaconnect_purchaseOfferingCmd)
}
