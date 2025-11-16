package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateThingCmd = &cobra.Command{
	Use:   "update-thing",
	Short: "Updates the data for a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateThingCmd).Standalone()

		iot_updateThingCmd.Flags().String("attribute-payload", "", "A list of thing attributes, a JSON string containing name-value pairs.")
		iot_updateThingCmd.Flags().String("expected-version", "", "The expected version of the thing record in the registry.")
		iot_updateThingCmd.Flags().String("remove-thing-type", "", "Remove a thing type association.")
		iot_updateThingCmd.Flags().String("thing-name", "", "The name of the thing to update.")
		iot_updateThingCmd.Flags().String("thing-type-name", "", "The name of the thing type.")
		iot_updateThingCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_updateThingCmd)
}
