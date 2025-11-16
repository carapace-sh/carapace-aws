package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deletePlaceIndexCmd = &cobra.Command{
	Use:   "delete-place-index",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deletePlaceIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_deletePlaceIndexCmd).Standalone()

		location_deletePlaceIndexCmd.Flags().String("index-name", "", "The name of the place index resource to be deleted.")
		location_deletePlaceIndexCmd.MarkFlagRequired("index-name")
	})
	locationCmd.AddCommand(location_deletePlaceIndexCmd)
}
