package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateThingGroupCmd = &cobra.Command{
	Use:   "update-thing-group",
	Short: "Update a thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateThingGroupCmd).Standalone()

		iot_updateThingGroupCmd.Flags().String("expected-version", "", "The expected version of the thing group.")
		iot_updateThingGroupCmd.Flags().String("thing-group-name", "", "The thing group to update.")
		iot_updateThingGroupCmd.Flags().String("thing-group-properties", "", "The thing group properties.")
		iot_updateThingGroupCmd.MarkFlagRequired("thing-group-name")
		iot_updateThingGroupCmd.MarkFlagRequired("thing-group-properties")
	})
	iotCmd.AddCommand(iot_updateThingGroupCmd)
}
