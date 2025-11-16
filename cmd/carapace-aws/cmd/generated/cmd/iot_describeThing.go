package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeThingCmd = &cobra.Command{
	Use:   "describe-thing",
	Short: "Gets information about the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeThingCmd).Standalone()

	iot_describeThingCmd.Flags().String("thing-name", "", "The name of the thing.")
	iot_describeThingCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_describeThingCmd)
}
