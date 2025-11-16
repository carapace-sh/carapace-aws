package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateThingTypeCmd = &cobra.Command{
	Use:   "update-thing-type",
	Short: "Updates a thing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateThingTypeCmd).Standalone()

	iot_updateThingTypeCmd.Flags().String("thing-type-name", "", "The name of a thing type.")
	iot_updateThingTypeCmd.Flags().String("thing-type-properties", "", "")
	iot_updateThingTypeCmd.MarkFlagRequired("thing-type-name")
	iotCmd.AddCommand(iot_updateThingTypeCmd)
}
