package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listGroundStationsCmd = &cobra.Command{
	Use:   "list-ground-stations",
	Short: "Returns a list of ground stations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listGroundStationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_listGroundStationsCmd).Standalone()

		groundstation_listGroundStationsCmd.Flags().String("max-results", "", "Maximum number of ground stations returned.")
		groundstation_listGroundStationsCmd.Flags().String("next-token", "", "Next token that can be supplied in the next call to get the next page of ground stations.")
		groundstation_listGroundStationsCmd.Flags().String("satellite-id", "", "Satellite ID to retrieve on-boarded ground stations.")
	})
	groundstationCmd.AddCommand(groundstation_listGroundStationsCmd)
}
