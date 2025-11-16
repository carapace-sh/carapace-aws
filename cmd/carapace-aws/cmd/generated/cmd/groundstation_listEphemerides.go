package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listEphemeridesCmd = &cobra.Command{
	Use:   "list-ephemerides",
	Short: "List your existing ephemerides.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listEphemeridesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_listEphemeridesCmd).Standalone()

		groundstation_listEphemeridesCmd.Flags().String("end-time", "", "The end time for the list operation in UTC.")
		groundstation_listEphemeridesCmd.Flags().String("ephemeris-type", "", "Filter ephemerides by type.")
		groundstation_listEphemeridesCmd.Flags().String("max-results", "", "Maximum number of ephemerides to return.")
		groundstation_listEphemeridesCmd.Flags().String("next-token", "", "Pagination token.")
		groundstation_listEphemeridesCmd.Flags().String("satellite-id", "", "The AWS Ground Station satellite ID to list ephemeris for.")
		groundstation_listEphemeridesCmd.Flags().String("start-time", "", "The start time for the list operation in UTC.")
		groundstation_listEphemeridesCmd.Flags().String("status-list", "", "The list of ephemeris status to return.")
		groundstation_listEphemeridesCmd.MarkFlagRequired("end-time")
		groundstation_listEphemeridesCmd.MarkFlagRequired("start-time")
	})
	groundstationCmd.AddCommand(groundstation_listEphemeridesCmd)
}
