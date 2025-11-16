package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_describeEphemerisCmd = &cobra.Command{
	Use:   "describe-ephemeris",
	Short: "Retrieve information about an existing ephemeris.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_describeEphemerisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_describeEphemerisCmd).Standalone()

		groundstation_describeEphemerisCmd.Flags().String("ephemeris-id", "", "The AWS Ground Station ephemeris ID.")
		groundstation_describeEphemerisCmd.MarkFlagRequired("ephemeris-id")
	})
	groundstationCmd.AddCommand(groundstation_describeEphemerisCmd)
}
