package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listOfferingsCmd = &cobra.Command{
	Use:   "list-offerings",
	Short: "Displays a list of all offerings that are available to this account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listOfferingsCmd).Standalone()

	mediaconnect_listOfferingsCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
	mediaconnect_listOfferingsCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	mediaconnectCmd.AddCommand(mediaconnect_listOfferingsCmd)
}
