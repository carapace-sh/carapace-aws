package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateCloudWatchAlarmTemplateCmd = &cobra.Command{
	Use:   "update-cloud-watch-alarm-template",
	Short: "Updates the specified cloudwatch alarm template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateCloudWatchAlarmTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateCloudWatchAlarmTemplateCmd).Standalone()

		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("comparison-operator", "", "")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("datapoints-to-alarm", "", "The number of datapoints within the evaluation period that must be breaching to trigger the alarm.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("description", "", "A resource's optional description.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("evaluation-periods", "", "The number of periods over which data is compared to the specified threshold.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("group-identifier", "", "A cloudwatch alarm template group's identifier.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("identifier", "", "A cloudwatch alarm template's identifier.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("metric-name", "", "The name of the metric associated with the alarm.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("name", "", "A resource's name.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("period", "", "The period, in seconds, over which the specified statistic is applied.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("statistic", "", "")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("target-resource-type", "", "")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("threshold", "", "The threshold value to compare with the specified statistic.")
		medialive_updateCloudWatchAlarmTemplateCmd.Flags().String("treat-missing-data", "", "")
		medialive_updateCloudWatchAlarmTemplateCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_updateCloudWatchAlarmTemplateCmd)
}
