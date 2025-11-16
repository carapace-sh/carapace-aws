package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_disableAlarmActionsCmd = &cobra.Command{
	Use:   "disable-alarm-actions",
	Short: "Disables the actions for the specified alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_disableAlarmActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_disableAlarmActionsCmd).Standalone()

		cloudwatch_disableAlarmActionsCmd.Flags().String("alarm-names", "", "The names of the alarms.")
		cloudwatch_disableAlarmActionsCmd.MarkFlagRequired("alarm-names")
	})
	cloudwatchCmd.AddCommand(cloudwatch_disableAlarmActionsCmd)
}
