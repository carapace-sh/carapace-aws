package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_purchaseOfferingCmd = &cobra.Command{
	Use:   "purchase-offering",
	Short: "Purchase an offering and create a reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_purchaseOfferingCmd).Standalone()

	medialive_purchaseOfferingCmd.Flags().String("count", "", "Number of resources")
	medialive_purchaseOfferingCmd.Flags().String("name", "", "Name for the new reservation")
	medialive_purchaseOfferingCmd.Flags().String("offering-id", "", "Offering to purchase, e.g. '87654321'")
	medialive_purchaseOfferingCmd.Flags().String("renewal-settings", "", "Renewal settings for the reservation")
	medialive_purchaseOfferingCmd.Flags().String("request-id", "", "Unique request ID to be specified.")
	medialive_purchaseOfferingCmd.Flags().String("start", "", "Requested reservation start time (UTC) in ISO-8601 format.")
	medialive_purchaseOfferingCmd.Flags().String("tags", "", "A collection of key-value pairs")
	medialive_purchaseOfferingCmd.MarkFlagRequired("count")
	medialive_purchaseOfferingCmd.MarkFlagRequired("offering-id")
	medialiveCmd.AddCommand(medialive_purchaseOfferingCmd)
}
