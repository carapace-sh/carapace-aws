package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_removeStreamGroupLocationsCmd = &cobra.Command{
	Use:   "remove-stream-group-locations",
	Short: "Removes a set of remote locations from this stream group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_removeStreamGroupLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_removeStreamGroupLocationsCmd).Standalone()

		gameliftstreams_removeStreamGroupLocationsCmd.Flags().String("identifier", "", "A stream group to remove the specified locations from.")
		gameliftstreams_removeStreamGroupLocationsCmd.Flags().String("locations", "", "A set of locations to remove this stream group.")
		gameliftstreams_removeStreamGroupLocationsCmd.MarkFlagRequired("identifier")
		gameliftstreams_removeStreamGroupLocationsCmd.MarkFlagRequired("locations")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_removeStreamGroupLocationsCmd)
}
