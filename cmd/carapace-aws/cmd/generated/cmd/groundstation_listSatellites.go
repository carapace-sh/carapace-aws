package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listSatellitesCmd = &cobra.Command{
	Use:   "list-satellites",
	Short: "Returns a list of satellites.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listSatellitesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_listSatellitesCmd).Standalone()

		groundstation_listSatellitesCmd.Flags().String("max-results", "", "Maximum number of satellites returned.")
		groundstation_listSatellitesCmd.Flags().String("next-token", "", "Next token that can be supplied in the next call to get the next page of satellites.")
	})
	groundstationCmd.AddCommand(groundstation_listSatellitesCmd)
}
