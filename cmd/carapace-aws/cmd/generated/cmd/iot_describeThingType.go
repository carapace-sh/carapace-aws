package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeThingTypeCmd = &cobra.Command{
	Use:   "describe-thing-type",
	Short: "Gets information about the specified thing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeThingTypeCmd).Standalone()

	iot_describeThingTypeCmd.Flags().String("thing-type-name", "", "The name of the thing type.")
	iot_describeThingTypeCmd.MarkFlagRequired("thing-type-name")
	iotCmd.AddCommand(iot_describeThingTypeCmd)
}
