package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createCloudWatchAlarmTemplateCmd = &cobra.Command{
	Use:   "create-cloud-watch-alarm-template",
	Short: "Creates a cloudwatch alarm template to dynamically generate cloudwatch metric alarms on targeted resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createCloudWatchAlarmTemplateCmd).Standalone()

	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("comparison-operator", "", "")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("datapoints-to-alarm", "", "The number of datapoints within the evaluation period that must be breaching to trigger the alarm.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("description", "", "A resource's optional description.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("evaluation-periods", "", "The number of periods over which data is compared to the specified threshold.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("group-identifier", "", "A cloudwatch alarm template group's identifier.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("metric-name", "", "The name of the metric associated with the alarm.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("name", "", "A resource's name.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("period", "", "The period, in seconds, over which the specified statistic is applied.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("statistic", "", "")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("tags", "", "")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("target-resource-type", "", "")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("threshold", "", "The threshold value to compare with the specified statistic.")
	medialive_createCloudWatchAlarmTemplateCmd.Flags().String("treat-missing-data", "", "")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("comparison-operator")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("evaluation-periods")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("group-identifier")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("metric-name")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("name")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("period")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("statistic")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("target-resource-type")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("threshold")
	medialive_createCloudWatchAlarmTemplateCmd.MarkFlagRequired("treat-missing-data")
	medialiveCmd.AddCommand(medialive_createCloudWatchAlarmTemplateCmd)
}
