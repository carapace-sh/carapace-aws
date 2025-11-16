package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listMissionProfilesCmd = &cobra.Command{
	Use:   "list-mission-profiles",
	Short: "Returns a list of mission profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listMissionProfilesCmd).Standalone()

	groundstation_listMissionProfilesCmd.Flags().String("max-results", "", "Maximum number of mission profiles returned.")
	groundstation_listMissionProfilesCmd.Flags().String("next-token", "", "Next token returned in the request of a previous `ListMissionProfiles` call.")
	groundstationCmd.AddCommand(groundstation_listMissionProfilesCmd)
}
