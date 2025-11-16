package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingGroupsForThingCmd = &cobra.Command{
	Use:   "list-thing-groups-for-thing",
	Short: "List the thing groups to which the specified thing belongs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingGroupsForThingCmd).Standalone()

	iot_listThingGroupsForThingCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listThingGroupsForThingCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listThingGroupsForThingCmd.Flags().String("thing-name", "", "The thing name.")
	iot_listThingGroupsForThingCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_listThingGroupsForThingCmd)
}
