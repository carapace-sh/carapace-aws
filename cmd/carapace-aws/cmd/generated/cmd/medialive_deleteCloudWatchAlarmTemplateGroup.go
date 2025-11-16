package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteCloudWatchAlarmTemplateGroupCmd = &cobra.Command{
	Use:   "delete-cloud-watch-alarm-template-group",
	Short: "Deletes a cloudwatch alarm template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteCloudWatchAlarmTemplateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteCloudWatchAlarmTemplateGroupCmd).Standalone()

		medialive_deleteCloudWatchAlarmTemplateGroupCmd.Flags().String("identifier", "", "A cloudwatch alarm template group's identifier.")
		medialive_deleteCloudWatchAlarmTemplateGroupCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_deleteCloudWatchAlarmTemplateGroupCmd)
}
