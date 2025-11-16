package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_updateEphemerisCmd = &cobra.Command{
	Use:   "update-ephemeris",
	Short: "Update an existing ephemeris.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_updateEphemerisCmd).Standalone()

	groundstation_updateEphemerisCmd.Flags().Bool("enabled", false, "Enable or disable the ephemeris.")
	groundstation_updateEphemerisCmd.Flags().String("ephemeris-id", "", "The AWS Ground Station ephemeris ID.")
	groundstation_updateEphemerisCmd.Flags().String("name", "", "A name that you can use to identify the ephemeris.")
	groundstation_updateEphemerisCmd.Flags().Bool("no-enabled", false, "Enable or disable the ephemeris.")
	groundstation_updateEphemerisCmd.Flags().String("priority", "", "A priority score that determines which ephemeris to use when multiple ephemerides overlap.")
	groundstation_updateEphemerisCmd.MarkFlagRequired("enabled")
	groundstation_updateEphemerisCmd.MarkFlagRequired("ephemeris-id")
	groundstation_updateEphemerisCmd.Flag("no-enabled").Hidden = true
	groundstation_updateEphemerisCmd.MarkFlagRequired("no-enabled")
	groundstationCmd.AddCommand(groundstation_updateEphemerisCmd)
}
