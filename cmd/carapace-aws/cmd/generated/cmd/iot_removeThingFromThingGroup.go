package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_removeThingFromThingGroupCmd = &cobra.Command{
	Use:   "remove-thing-from-thing-group",
	Short: "Remove the specified thing from the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_removeThingFromThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_removeThingFromThingGroupCmd).Standalone()

		iot_removeThingFromThingGroupCmd.Flags().String("thing-arn", "", "The ARN of the thing to remove from the group.")
		iot_removeThingFromThingGroupCmd.Flags().String("thing-group-arn", "", "The group ARN.")
		iot_removeThingFromThingGroupCmd.Flags().String("thing-group-name", "", "The group name.")
		iot_removeThingFromThingGroupCmd.Flags().String("thing-name", "", "The name of the thing to remove from the group.")
	})
	iotCmd.AddCommand(iot_removeThingFromThingGroupCmd)
}
