package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_addThingToThingGroupCmd = &cobra.Command{
	Use:   "add-thing-to-thing-group",
	Short: "Adds a thing to a thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_addThingToThingGroupCmd).Standalone()

	iot_addThingToThingGroupCmd.Flags().String("override-dynamic-groups", "", "Override dynamic thing groups with static thing groups when 10-group limit is reached.")
	iot_addThingToThingGroupCmd.Flags().String("thing-arn", "", "The ARN of the thing to add to a group.")
	iot_addThingToThingGroupCmd.Flags().String("thing-group-arn", "", "The ARN of the group to which you are adding a thing.")
	iot_addThingToThingGroupCmd.Flags().String("thing-group-name", "", "The name of the group to which you are adding a thing.")
	iot_addThingToThingGroupCmd.Flags().String("thing-name", "", "The name of the thing to add to a group.")
	iotCmd.AddCommand(iot_addThingToThingGroupCmd)
}
