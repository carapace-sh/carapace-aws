package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_testAlarmCmd = &cobra.Command{
	Use:   "test-alarm",
	Short: "Tests an alarm by displaying a banner on the Amazon Lightsail console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_testAlarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_testAlarmCmd).Standalone()

		lightsail_testAlarmCmd.Flags().String("alarm-name", "", "The name of the alarm to test.")
		lightsail_testAlarmCmd.Flags().String("state", "", "The alarm state to test.")
		lightsail_testAlarmCmd.MarkFlagRequired("alarm-name")
		lightsail_testAlarmCmd.MarkFlagRequired("state")
	})
	lightsailCmd.AddCommand(lightsail_testAlarmCmd)
}
