package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeThingGroupCmd = &cobra.Command{
	Use:   "describe-thing-group",
	Short: "Describe a thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeThingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeThingGroupCmd).Standalone()

		iot_describeThingGroupCmd.Flags().String("thing-group-name", "", "The name of the thing group.")
		iot_describeThingGroupCmd.MarkFlagRequired("thing-group-name")
	})
	iotCmd.AddCommand(iot_describeThingGroupCmd)
}
