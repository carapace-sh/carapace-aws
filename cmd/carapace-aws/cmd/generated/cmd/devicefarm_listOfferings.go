package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listOfferingsCmd = &cobra.Command{
	Use:   "list-offerings",
	Short: "Returns a list of products or offerings that the user can manage through the API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listOfferingsCmd).Standalone()

	devicefarm_listOfferingsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarmCmd.AddCommand(devicefarm_listOfferingsCmd)
}
