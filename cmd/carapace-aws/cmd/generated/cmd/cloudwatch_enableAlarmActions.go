package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_enableAlarmActionsCmd = &cobra.Command{
	Use:   "enable-alarm-actions",
	Short: "Enables the actions for the specified alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_enableAlarmActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_enableAlarmActionsCmd).Standalone()

		cloudwatch_enableAlarmActionsCmd.Flags().String("alarm-names", "", "The names of the alarms.")
		cloudwatch_enableAlarmActionsCmd.MarkFlagRequired("alarm-names")
	})
	cloudwatchCmd.AddCommand(cloudwatch_enableAlarmActionsCmd)
}
