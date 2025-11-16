package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_putMaintenanceStartTimeCmd = &cobra.Command{
	Use:   "put-maintenance-start-time",
	Short: "Set the maintenance start time for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_putMaintenanceStartTimeCmd).Standalone()

	backupGateway_putMaintenanceStartTimeCmd.Flags().String("day-of-month", "", "The day of the month start maintenance on a gateway.")
	backupGateway_putMaintenanceStartTimeCmd.Flags().String("day-of-week", "", "The day of the week to start maintenance on a gateway.")
	backupGateway_putMaintenanceStartTimeCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) for the gateway, used to specify its maintenance start time.")
	backupGateway_putMaintenanceStartTimeCmd.Flags().String("hour-of-day", "", "The hour of the day to start maintenance on a gateway.")
	backupGateway_putMaintenanceStartTimeCmd.Flags().String("minute-of-hour", "", "The minute of the hour to start maintenance on a gateway.")
	backupGateway_putMaintenanceStartTimeCmd.MarkFlagRequired("gateway-arn")
	backupGateway_putMaintenanceStartTimeCmd.MarkFlagRequired("hour-of-day")
	backupGateway_putMaintenanceStartTimeCmd.MarkFlagRequired("minute-of-hour")
	backupGatewayCmd.AddCommand(backupGateway_putMaintenanceStartTimeCmd)
}
