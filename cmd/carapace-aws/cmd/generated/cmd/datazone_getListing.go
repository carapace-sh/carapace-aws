package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getListingCmd = &cobra.Command{
	Use:   "get-listing",
	Short: "Gets a listing (a record of an asset at a given time).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getListingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getListingCmd).Standalone()

		datazone_getListingCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain.")
		datazone_getListingCmd.Flags().String("identifier", "", "The ID of the listing.")
		datazone_getListingCmd.Flags().String("listing-revision", "", "The revision of the listing.")
		datazone_getListingCmd.MarkFlagRequired("domain-identifier")
		datazone_getListingCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getListingCmd)
}
