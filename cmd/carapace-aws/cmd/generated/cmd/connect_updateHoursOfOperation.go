package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateHoursOfOperationCmd = &cobra.Command{
	Use:   "update-hours-of-operation",
	Short: "Updates the hours of operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateHoursOfOperationCmd).Standalone()

	connect_updateHoursOfOperationCmd.Flags().String("config", "", "Configuration information of the hours of operation.")
	connect_updateHoursOfOperationCmd.Flags().String("description", "", "The description of the hours of operation.")
	connect_updateHoursOfOperationCmd.Flags().String("hours-of-operation-id", "", "The identifier of the hours of operation.")
	connect_updateHoursOfOperationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateHoursOfOperationCmd.Flags().String("name", "", "The name of the hours of operation.")
	connect_updateHoursOfOperationCmd.Flags().String("time-zone", "", "The time zone of the hours of operation.")
	connect_updateHoursOfOperationCmd.MarkFlagRequired("hours-of-operation-id")
	connect_updateHoursOfOperationCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateHoursOfOperationCmd)
}
