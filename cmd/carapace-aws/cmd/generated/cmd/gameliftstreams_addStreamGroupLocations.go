package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_addStreamGroupLocationsCmd = &cobra.Command{
	Use:   "add-stream-group-locations",
	Short: "Add locations that can host stream sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_addStreamGroupLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_addStreamGroupLocationsCmd).Standalone()

		gameliftstreams_addStreamGroupLocationsCmd.Flags().String("identifier", "", "A stream group to add the specified locations to.")
		gameliftstreams_addStreamGroupLocationsCmd.Flags().String("location-configurations", "", "A set of one or more locations and the streaming capacity for each location.")
		gameliftstreams_addStreamGroupLocationsCmd.MarkFlagRequired("identifier")
		gameliftstreams_addStreamGroupLocationsCmd.MarkFlagRequired("location-configurations")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_addStreamGroupLocationsCmd)
}
