package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createThingCmd = &cobra.Command{
	Use:   "create-thing",
	Short: "Creates a thing record in the registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createThingCmd).Standalone()

		iot_createThingCmd.Flags().String("attribute-payload", "", "The attribute payload, which consists of up to three name/value pairs in a JSON document.")
		iot_createThingCmd.Flags().String("billing-group-name", "", "The name of the billing group the thing will be added to.")
		iot_createThingCmd.Flags().String("thing-name", "", "The name of the thing to create.")
		iot_createThingCmd.Flags().String("thing-type-name", "", "The name of the thing type associated with the new thing.")
		iot_createThingCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_createThingCmd)
}
