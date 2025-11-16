package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_deleteEphemerisCmd = &cobra.Command{
	Use:   "delete-ephemeris",
	Short: "Delete an ephemeris.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_deleteEphemerisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_deleteEphemerisCmd).Standalone()

		groundstation_deleteEphemerisCmd.Flags().String("ephemeris-id", "", "The AWS Ground Station ephemeris ID.")
		groundstation_deleteEphemerisCmd.MarkFlagRequired("ephemeris-id")
	})
	groundstationCmd.AddCommand(groundstation_deleteEphemerisCmd)
}
