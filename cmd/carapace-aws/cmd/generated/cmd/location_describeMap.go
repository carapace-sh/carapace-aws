package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describeMapCmd = &cobra.Command{
	Use:   "describe-map",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describeMapCmd).Standalone()

	location_describeMapCmd.Flags().String("map-name", "", "The name of the map resource.")
	location_describeMapCmd.MarkFlagRequired("map-name")
	locationCmd.AddCommand(location_describeMapCmd)
}
