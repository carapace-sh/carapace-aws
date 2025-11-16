package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteThingCmd = &cobra.Command{
	Use:   "delete-thing",
	Short: "Deletes the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteThingCmd).Standalone()

	iot_deleteThingCmd.Flags().String("expected-version", "", "The expected version of the thing record in the registry.")
	iot_deleteThingCmd.Flags().String("thing-name", "", "The name of the thing to delete.")
	iot_deleteThingCmd.MarkFlagRequired("thing-name")
	iotCmd.AddCommand(iot_deleteThingCmd)
}
