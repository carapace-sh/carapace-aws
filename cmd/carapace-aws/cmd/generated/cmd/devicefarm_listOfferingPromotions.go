package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listOfferingPromotionsCmd = &cobra.Command{
	Use:   "list-offering-promotions",
	Short: "Returns a list of offering promotions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listOfferingPromotionsCmd).Standalone()

	devicefarm_listOfferingPromotionsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarmCmd.AddCommand(devicefarm_listOfferingPromotionsCmd)
}
