package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateThingGroupsForThingCmd = &cobra.Command{
	Use:   "update-thing-groups-for-thing",
	Short: "Updates the groups to which the thing belongs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateThingGroupsForThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateThingGroupsForThingCmd).Standalone()

		iot_updateThingGroupsForThingCmd.Flags().String("override-dynamic-groups", "", "Override dynamic thing groups with static thing groups when 10-group limit is reached.")
		iot_updateThingGroupsForThingCmd.Flags().String("thing-groups-to-add", "", "The groups to which the thing will be added.")
		iot_updateThingGroupsForThingCmd.Flags().String("thing-groups-to-remove", "", "The groups from which the thing will be removed.")
		iot_updateThingGroupsForThingCmd.Flags().String("thing-name", "", "The thing whose group memberships will be updated.")
	})
	iotCmd.AddCommand(iot_updateThingGroupsForThingCmd)
}
