package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_getCloudWatchAlarmTemplateGroupCmd = &cobra.Command{
	Use:   "get-cloud-watch-alarm-template-group",
	Short: "Retrieves the specified cloudwatch alarm template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_getCloudWatchAlarmTemplateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_getCloudWatchAlarmTemplateGroupCmd).Standalone()

		medialive_getCloudWatchAlarmTemplateGroupCmd.Flags().String("identifier", "", "A cloudwatch alarm template group's identifier.")
		medialive_getCloudWatchAlarmTemplateGroupCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_getCloudWatchAlarmTemplateGroupCmd)
}
