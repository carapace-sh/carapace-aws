package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteThingTypeCmd = &cobra.Command{
	Use:   "delete-thing-type",
	Short: "Deletes the specified thing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteThingTypeCmd).Standalone()

	iot_deleteThingTypeCmd.Flags().String("thing-type-name", "", "The name of the thing type.")
	iot_deleteThingTypeCmd.MarkFlagRequired("thing-type-name")
	iotCmd.AddCommand(iot_deleteThingTypeCmd)
}
