package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateHoursOfOperationOverrideCmd = &cobra.Command{
	Use:   "update-hours-of-operation-override",
	Short: "Update the hours of operation override.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateHoursOfOperationOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateHoursOfOperationOverrideCmd).Standalone()

		connect_updateHoursOfOperationOverrideCmd.Flags().String("config", "", "Configuration information for the hours of operation override: day, start time, and end time.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("description", "", "The description of the hours of operation override.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("effective-from", "", "The date from when the hours of operation override would be effective.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("effective-till", "", "The date until the hours of operation override is effective.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-override-id", "", "The identifier for the hours of operation override.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateHoursOfOperationOverrideCmd.Flags().String("name", "", "The name of the hours of operation override.")
		connect_updateHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-id")
		connect_updateHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-override-id")
		connect_updateHoursOfOperationOverrideCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateHoursOfOperationOverrideCmd)
}
