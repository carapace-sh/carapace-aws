package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteDynamicThingGroupCmd = &cobra.Command{
	Use:   "delete-dynamic-thing-group",
	Short: "Deletes a dynamic thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteDynamicThingGroupCmd).Standalone()

	iot_deleteDynamicThingGroupCmd.Flags().String("expected-version", "", "The expected version of the dynamic thing group to delete.")
	iot_deleteDynamicThingGroupCmd.Flags().String("thing-group-name", "", "The name of the dynamic thing group to delete.")
	iot_deleteDynamicThingGroupCmd.MarkFlagRequired("thing-group-name")
	iotCmd.AddCommand(iot_deleteDynamicThingGroupCmd)
}
