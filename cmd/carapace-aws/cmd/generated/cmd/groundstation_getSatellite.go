package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getSatelliteCmd = &cobra.Command{
	Use:   "get-satellite",
	Short: "Returns a satellite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getSatelliteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_getSatelliteCmd).Standalone()

		groundstation_getSatelliteCmd.Flags().String("satellite-id", "", "UUID of a satellite.")
		groundstation_getSatelliteCmd.MarkFlagRequired("satellite-id")
	})
	groundstationCmd.AddCommand(groundstation_getSatelliteCmd)
}
