package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeAlarmContributorsCmd = &cobra.Command{
	Use:   "describe-alarm-contributors",
	Short: "Returns the information of the current alarm contributors that are in `ALARM` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeAlarmContributorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_describeAlarmContributorsCmd).Standalone()

		cloudwatch_describeAlarmContributorsCmd.Flags().String("alarm-name", "", "The name of the alarm for which to retrieve contributor information.")
		cloudwatch_describeAlarmContributorsCmd.Flags().String("next-token", "", "The token returned by a previous call to indicate that there is more data available.")
		cloudwatch_describeAlarmContributorsCmd.MarkFlagRequired("alarm-name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_describeAlarmContributorsCmd)
}
