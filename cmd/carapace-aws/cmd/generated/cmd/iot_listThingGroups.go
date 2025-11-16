package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingGroupsCmd = &cobra.Command{
	Use:   "list-thing-groups",
	Short: "List the thing groups in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listThingGroupsCmd).Standalone()

		iot_listThingGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listThingGroupsCmd.Flags().String("name-prefix-filter", "", "A filter that limits the results to those with the specified name prefix.")
		iot_listThingGroupsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listThingGroupsCmd.Flags().String("parent-group", "", "A filter that limits the results to those with the specified parent group.")
		iot_listThingGroupsCmd.Flags().String("recursive", "", "If true, return child groups as well.")
	})
	iotCmd.AddCommand(iot_listThingGroupsCmd)
}
