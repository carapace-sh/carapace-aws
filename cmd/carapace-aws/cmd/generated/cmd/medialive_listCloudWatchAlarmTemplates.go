package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listCloudWatchAlarmTemplatesCmd = &cobra.Command{
	Use:   "list-cloud-watch-alarm-templates",
	Short: "Lists cloudwatch alarm templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listCloudWatchAlarmTemplatesCmd).Standalone()

	medialive_listCloudWatchAlarmTemplatesCmd.Flags().String("group-identifier", "", "A cloudwatch alarm template group's identifier.")
	medialive_listCloudWatchAlarmTemplatesCmd.Flags().String("max-results", "", "")
	medialive_listCloudWatchAlarmTemplatesCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results in paginated list responses.")
	medialive_listCloudWatchAlarmTemplatesCmd.Flags().String("scope", "", "Represents the scope of a resource, with options for all scopes, AWS provided resources, or local resources.")
	medialive_listCloudWatchAlarmTemplatesCmd.Flags().String("signal-map-identifier", "", "A signal map's identifier.")
	medialiveCmd.AddCommand(medialive_listCloudWatchAlarmTemplatesCmd)
}
