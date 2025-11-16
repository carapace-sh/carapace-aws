package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listCloudWatchAlarmTemplateGroupsCmd = &cobra.Command{
	Use:   "list-cloud-watch-alarm-template-groups",
	Short: "Lists cloudwatch alarm template groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listCloudWatchAlarmTemplateGroupsCmd).Standalone()

	medialive_listCloudWatchAlarmTemplateGroupsCmd.Flags().String("max-results", "", "")
	medialive_listCloudWatchAlarmTemplateGroupsCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results in paginated list responses.")
	medialive_listCloudWatchAlarmTemplateGroupsCmd.Flags().String("scope", "", "Represents the scope of a resource, with options for all scopes, AWS provided resources, or local resources.")
	medialive_listCloudWatchAlarmTemplateGroupsCmd.Flags().String("signal-map-identifier", "", "A signal map's identifier.")
	medialiveCmd.AddCommand(medialive_listCloudWatchAlarmTemplateGroupsCmd)
}
