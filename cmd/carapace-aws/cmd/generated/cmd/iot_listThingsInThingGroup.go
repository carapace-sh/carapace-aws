package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingsInThingGroupCmd = &cobra.Command{
	Use:   "list-things-in-thing-group",
	Short: "Lists the things in the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingsInThingGroupCmd).Standalone()

	iot_listThingsInThingGroupCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listThingsInThingGroupCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listThingsInThingGroupCmd.Flags().String("recursive", "", "When true, list things in this thing group and in all child groups as well.")
	iot_listThingsInThingGroupCmd.Flags().String("thing-group-name", "", "The thing group name.")
	iot_listThingsInThingGroupCmd.MarkFlagRequired("thing-group-name")
	iotCmd.AddCommand(iot_listThingsInThingGroupCmd)
}
