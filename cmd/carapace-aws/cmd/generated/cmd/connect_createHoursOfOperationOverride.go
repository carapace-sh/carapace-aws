package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createHoursOfOperationOverrideCmd = &cobra.Command{
	Use:   "create-hours-of-operation-override",
	Short: "Creates an hours of operation override in an Amazon Connect hours of operation resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createHoursOfOperationOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createHoursOfOperationOverrideCmd).Standalone()

		connect_createHoursOfOperationOverrideCmd.Flags().String("config", "", "Configuration information for the hours of operation override: day, start time, and end time.")
		connect_createHoursOfOperationOverrideCmd.Flags().String("description", "", "The description of the hours of operation override.")
		connect_createHoursOfOperationOverrideCmd.Flags().String("effective-from", "", "The date from when the hours of operation override is effective.")
		connect_createHoursOfOperationOverrideCmd.Flags().String("effective-till", "", "The date until when the hours of operation override is effective.")
		connect_createHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation")
		connect_createHoursOfOperationOverrideCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createHoursOfOperationOverrideCmd.Flags().String("name", "", "The name of the hours of operation override.")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("config")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("effective-from")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("effective-till")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-id")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("instance-id")
		connect_createHoursOfOperationOverrideCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_createHoursOfOperationOverrideCmd)
}
