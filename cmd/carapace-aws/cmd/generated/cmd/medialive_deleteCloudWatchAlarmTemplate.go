package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteCloudWatchAlarmTemplateCmd = &cobra.Command{
	Use:   "delete-cloud-watch-alarm-template",
	Short: "Deletes a cloudwatch alarm template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteCloudWatchAlarmTemplateCmd).Standalone()

	medialive_deleteCloudWatchAlarmTemplateCmd.Flags().String("identifier", "", "A cloudwatch alarm template's identifier.")
	medialive_deleteCloudWatchAlarmTemplateCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_deleteCloudWatchAlarmTemplateCmd)
}
