package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteListingCmd = &cobra.Command{
	Use:   "delete-listing",
	Short: "Deletes a listing (a record of an asset at a given time).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteListingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteListingCmd).Standalone()

		datazone_deleteListingCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain.")
		datazone_deleteListingCmd.Flags().String("identifier", "", "The ID of the listing to be deleted.")
		datazone_deleteListingCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteListingCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteListingCmd)
}
