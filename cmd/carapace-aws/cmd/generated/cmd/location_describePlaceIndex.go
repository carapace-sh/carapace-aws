package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describePlaceIndexCmd = &cobra.Command{
	Use:   "describe-place-index",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describePlaceIndexCmd).Standalone()

	location_describePlaceIndexCmd.Flags().String("index-name", "", "The name of the place index resource.")
	location_describePlaceIndexCmd.MarkFlagRequired("index-name")
	locationCmd.AddCommand(location_describePlaceIndexCmd)
}
