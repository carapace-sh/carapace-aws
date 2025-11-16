package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createHoursOfOperationCmd = &cobra.Command{
	Use:   "create-hours-of-operation",
	Short: "Creates hours of operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createHoursOfOperationCmd).Standalone()

	connect_createHoursOfOperationCmd.Flags().String("config", "", "Configuration information for the hours of operation: day, start time, and end time.")
	connect_createHoursOfOperationCmd.Flags().String("description", "", "The description of the hours of operation.")
	connect_createHoursOfOperationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createHoursOfOperationCmd.Flags().String("name", "", "The name of the hours of operation.")
	connect_createHoursOfOperationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createHoursOfOperationCmd.Flags().String("time-zone", "", "The time zone of the hours of operation.")
	connect_createHoursOfOperationCmd.MarkFlagRequired("config")
	connect_createHoursOfOperationCmd.MarkFlagRequired("instance-id")
	connect_createHoursOfOperationCmd.MarkFlagRequired("name")
	connect_createHoursOfOperationCmd.MarkFlagRequired("time-zone")
	connectCmd.AddCommand(connect_createHoursOfOperationCmd)
}
