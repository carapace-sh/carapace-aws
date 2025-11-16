package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createThingGroupCmd = &cobra.Command{
	Use:   "create-thing-group",
	Short: "Create a thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createThingGroupCmd).Standalone()

	iot_createThingGroupCmd.Flags().String("parent-group-name", "", "The name of the parent thing group.")
	iot_createThingGroupCmd.Flags().String("tags", "", "Metadata which can be used to manage the thing group.")
	iot_createThingGroupCmd.Flags().String("thing-group-name", "", "The thing group name to create.")
	iot_createThingGroupCmd.Flags().String("thing-group-properties", "", "The thing group properties.")
	iot_createThingGroupCmd.MarkFlagRequired("thing-group-name")
	iotCmd.AddCommand(iot_createThingGroupCmd)
}
