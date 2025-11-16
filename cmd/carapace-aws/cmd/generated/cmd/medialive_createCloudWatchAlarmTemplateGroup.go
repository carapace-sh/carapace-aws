package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createCloudWatchAlarmTemplateGroupCmd = &cobra.Command{
	Use:   "create-cloud-watch-alarm-template-group",
	Short: "Creates a cloudwatch alarm template group to group your cloudwatch alarm templates and to attach to signal maps for dynamically creating alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createCloudWatchAlarmTemplateGroupCmd).Standalone()

	medialive_createCloudWatchAlarmTemplateGroupCmd.Flags().String("description", "", "A resource's optional description.")
	medialive_createCloudWatchAlarmTemplateGroupCmd.Flags().String("name", "", "A resource's name.")
	medialive_createCloudWatchAlarmTemplateGroupCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
	medialive_createCloudWatchAlarmTemplateGroupCmd.Flags().String("tags", "", "")
	medialive_createCloudWatchAlarmTemplateGroupCmd.MarkFlagRequired("name")
	medialiveCmd.AddCommand(medialive_createCloudWatchAlarmTemplateGroupCmd)
}
