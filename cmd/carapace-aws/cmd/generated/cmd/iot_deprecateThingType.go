package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deprecateThingTypeCmd = &cobra.Command{
	Use:   "deprecate-thing-type",
	Short: "Deprecates a thing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deprecateThingTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deprecateThingTypeCmd).Standalone()

		iot_deprecateThingTypeCmd.Flags().String("thing-type-name", "", "The name of the thing type to deprecate.")
		iot_deprecateThingTypeCmd.Flags().String("undo-deprecate", "", "Whether to undeprecate a deprecated thing type.")
		iot_deprecateThingTypeCmd.MarkFlagRequired("thing-type-name")
	})
	iotCmd.AddCommand(iot_deprecateThingTypeCmd)
}
