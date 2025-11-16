package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getEffectiveHoursOfOperationsCmd = &cobra.Command{
	Use:   "get-effective-hours-of-operations",
	Short: "Get the hours of operations with the effective override applied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getEffectiveHoursOfOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getEffectiveHoursOfOperationsCmd).Standalone()

		connect_getEffectiveHoursOfOperationsCmd.Flags().String("from-date", "", "The date from when the hours of operation are listed.")
		connect_getEffectiveHoursOfOperationsCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_getEffectiveHoursOfOperationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getEffectiveHoursOfOperationsCmd.Flags().String("to-date", "", "The date until when the hours of operation are listed.")
		connect_getEffectiveHoursOfOperationsCmd.MarkFlagRequired("from-date")
		connect_getEffectiveHoursOfOperationsCmd.MarkFlagRequired("hours-of-operation-id")
		connect_getEffectiveHoursOfOperationsCmd.MarkFlagRequired("instance-id")
		connect_getEffectiveHoursOfOperationsCmd.MarkFlagRequired("to-date")
	})
	connectCmd.AddCommand(connect_getEffectiveHoursOfOperationsCmd)
}
