package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_createEphemerisCmd = &cobra.Command{
	Use:   "create-ephemeris",
	Short: "Create an ephemeris with your specified [EphemerisData]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_createEphemerisCmd).Standalone()

	groundstation_createEphemerisCmd.Flags().Bool("enabled", false, "Set to `true` to enable the ephemeris after validation.")
	groundstation_createEphemerisCmd.Flags().String("ephemeris", "", "Ephemeris data.")
	groundstation_createEphemerisCmd.Flags().String("expiration-time", "", "An overall expiration time for the ephemeris in UTC, after which it will become `EXPIRED`.")
	groundstation_createEphemerisCmd.Flags().String("kms-key-arn", "", "The ARN of the KMS key to use for encrypting the ephemeris.")
	groundstation_createEphemerisCmd.Flags().String("name", "", "A name that you can use to identify the ephemeris.")
	groundstation_createEphemerisCmd.Flags().Bool("no-enabled", false, "Set to `true` to enable the ephemeris after validation.")
	groundstation_createEphemerisCmd.Flags().String("priority", "", "A priority score that determines which ephemeris to use when multiple ephemerides overlap.")
	groundstation_createEphemerisCmd.Flags().String("satellite-id", "", "The satellite ID that associates this ephemeris with a satellite in AWS Ground Station.")
	groundstation_createEphemerisCmd.Flags().String("tags", "", "Tags assigned to an ephemeris.")
	groundstation_createEphemerisCmd.MarkFlagRequired("name")
	groundstation_createEphemerisCmd.Flag("no-enabled").Hidden = true
	groundstationCmd.AddCommand(groundstation_createEphemerisCmd)
}
