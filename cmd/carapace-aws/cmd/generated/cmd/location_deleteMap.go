package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deleteMapCmd = &cobra.Command{
	Use:   "delete-map",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deleteMapCmd).Standalone()

	location_deleteMapCmd.Flags().String("map-name", "", "The name of the map resource to be deleted.")
	location_deleteMapCmd.MarkFlagRequired("map-name")
	locationCmd.AddCommand(location_deleteMapCmd)
}
