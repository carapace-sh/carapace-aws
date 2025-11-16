package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createThingTypeCmd = &cobra.Command{
	Use:   "create-thing-type",
	Short: "Creates a new thing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createThingTypeCmd).Standalone()

	iot_createThingTypeCmd.Flags().String("tags", "", "Metadata which can be used to manage the thing type.")
	iot_createThingTypeCmd.Flags().String("thing-type-name", "", "The name of the thing type.")
	iot_createThingTypeCmd.Flags().String("thing-type-properties", "", "The ThingTypeProperties for the thing type to create.")
	iot_createThingTypeCmd.MarkFlagRequired("thing-type-name")
	iotCmd.AddCommand(iot_createThingTypeCmd)
}
