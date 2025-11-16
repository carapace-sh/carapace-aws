package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createDynamicThingGroupCmd = &cobra.Command{
	Use:   "create-dynamic-thing-group",
	Short: "Creates a dynamic thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createDynamicThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createDynamicThingGroupCmd).Standalone()

		iot_createDynamicThingGroupCmd.Flags().String("index-name", "", "The dynamic thing group index name.")
		iot_createDynamicThingGroupCmd.Flags().String("query-string", "", "The dynamic thing group search query string.")
		iot_createDynamicThingGroupCmd.Flags().String("query-version", "", "The dynamic thing group query version.")
		iot_createDynamicThingGroupCmd.Flags().String("tags", "", "Metadata which can be used to manage the dynamic thing group.")
		iot_createDynamicThingGroupCmd.Flags().String("thing-group-name", "", "The dynamic thing group name to create.")
		iot_createDynamicThingGroupCmd.Flags().String("thing-group-properties", "", "The dynamic thing group properties.")
		iot_createDynamicThingGroupCmd.MarkFlagRequired("query-string")
		iot_createDynamicThingGroupCmd.MarkFlagRequired("thing-group-name")
	})
	iotCmd.AddCommand(iot_createDynamicThingGroupCmd)
}
