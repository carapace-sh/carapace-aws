package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getMinuteUsageCmd = &cobra.Command{
	Use:   "get-minute-usage",
	Short: "Returns the number of reserved minutes used by account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getMinuteUsageCmd).Standalone()

	groundstation_getMinuteUsageCmd.Flags().String("month", "", "The month being requested, with a value of 1-12.")
	groundstation_getMinuteUsageCmd.Flags().String("year", "", "The year being requested, in the format of YYYY.")
	groundstation_getMinuteUsageCmd.MarkFlagRequired("month")
	groundstation_getMinuteUsageCmd.MarkFlagRequired("year")
	groundstationCmd.AddCommand(groundstation_getMinuteUsageCmd)
}
