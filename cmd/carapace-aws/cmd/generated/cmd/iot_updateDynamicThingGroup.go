package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateDynamicThingGroupCmd = &cobra.Command{
	Use:   "update-dynamic-thing-group",
	Short: "Updates a dynamic thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateDynamicThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateDynamicThingGroupCmd).Standalone()

		iot_updateDynamicThingGroupCmd.Flags().String("expected-version", "", "The expected version of the dynamic thing group to update.")
		iot_updateDynamicThingGroupCmd.Flags().String("index-name", "", "The dynamic thing group index to update.")
		iot_updateDynamicThingGroupCmd.Flags().String("query-string", "", "The dynamic thing group search query string to update.")
		iot_updateDynamicThingGroupCmd.Flags().String("query-version", "", "The dynamic thing group query version to update.")
		iot_updateDynamicThingGroupCmd.Flags().String("thing-group-name", "", "The name of the dynamic thing group to update.")
		iot_updateDynamicThingGroupCmd.Flags().String("thing-group-properties", "", "The dynamic thing group properties to update.")
		iot_updateDynamicThingGroupCmd.MarkFlagRequired("thing-group-name")
		iot_updateDynamicThingGroupCmd.MarkFlagRequired("thing-group-properties")
	})
	iotCmd.AddCommand(iot_updateDynamicThingGroupCmd)
}
