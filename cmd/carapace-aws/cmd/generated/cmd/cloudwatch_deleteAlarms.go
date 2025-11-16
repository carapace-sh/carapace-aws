package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_deleteAlarmsCmd = &cobra.Command{
	Use:   "delete-alarms",
	Short: "Deletes the specified alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_deleteAlarmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_deleteAlarmsCmd).Standalone()

		cloudwatch_deleteAlarmsCmd.Flags().String("alarm-names", "", "The alarms to be deleted.")
		cloudwatch_deleteAlarmsCmd.MarkFlagRequired("alarm-names")
	})
	cloudwatchCmd.AddCommand(cloudwatch_deleteAlarmsCmd)
}
