package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_getCloudWatchAlarmTemplateCmd = &cobra.Command{
	Use:   "get-cloud-watch-alarm-template",
	Short: "Retrieves the specified cloudwatch alarm template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_getCloudWatchAlarmTemplateCmd).Standalone()

	medialive_getCloudWatchAlarmTemplateCmd.Flags().String("identifier", "", "A cloudwatch alarm template's identifier.")
	medialive_getCloudWatchAlarmTemplateCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_getCloudWatchAlarmTemplateCmd)
}
