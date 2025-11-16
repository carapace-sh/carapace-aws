package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateMaintenanceStartTimeCmd = &cobra.Command{
	Use:   "update-maintenance-start-time",
	Short: "Updates a gateway's maintenance window schedule, with settings for monthly or weekly cadence, specific day and time to begin maintenance, and which types of updates to apply.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateMaintenanceStartTimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateMaintenanceStartTimeCmd).Standalone()

		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("day-of-month", "", "The day of the month component of the maintenance start time represented as an ordinal number from 1 to 28, where 1 represents the first day of the month.")
		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("day-of-week", "", "The day of the week component of the maintenance start time week represented as an ordinal number from 0 to 6, where 0 represents Sunday and 6 represents Saturday.")
		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("gateway-arn", "", "")
		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("hour-of-day", "", "The hour component of the maintenance start time represented as *hh*, where *hh* is the hour (00 to 23).")
		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("minute-of-hour", "", "The minute component of the maintenance start time represented as *mm*, where *mm* is the minute (00 to 59).")
		storagegateway_updateMaintenanceStartTimeCmd.Flags().String("software-update-preferences", "", "A set of variables indicating the software update preferences for the gateway.")
		storagegateway_updateMaintenanceStartTimeCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateMaintenanceStartTimeCmd)
}
