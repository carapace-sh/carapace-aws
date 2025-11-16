package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteThingGroupCmd = &cobra.Command{
	Use:   "delete-thing-group",
	Short: "Deletes a thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteThingGroupCmd).Standalone()

		iot_deleteThingGroupCmd.Flags().String("expected-version", "", "The expected version of the thing group to delete.")
		iot_deleteThingGroupCmd.Flags().String("thing-group-name", "", "The name of the thing group to delete.")
		iot_deleteThingGroupCmd.MarkFlagRequired("thing-group-name")
	})
	iotCmd.AddCommand(iot_deleteThingGroupCmd)
}
