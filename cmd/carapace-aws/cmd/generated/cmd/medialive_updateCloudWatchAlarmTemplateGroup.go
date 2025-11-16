package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateCloudWatchAlarmTemplateGroupCmd = &cobra.Command{
	Use:   "update-cloud-watch-alarm-template-group",
	Short: "Updates the specified cloudwatch alarm template group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateCloudWatchAlarmTemplateGroupCmd).Standalone()

	medialive_updateCloudWatchAlarmTemplateGroupCmd.Flags().String("description", "", "A resource's optional description.")
	medialive_updateCloudWatchAlarmTemplateGroupCmd.Flags().String("identifier", "", "A cloudwatch alarm template group's identifier.")
	medialive_updateCloudWatchAlarmTemplateGroupCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_updateCloudWatchAlarmTemplateGroupCmd)
}
