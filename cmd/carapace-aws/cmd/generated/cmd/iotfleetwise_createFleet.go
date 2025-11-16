package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "Creates a fleet that represents a group of vehicles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createFleetCmd).Standalone()

	iotfleetwise_createFleetCmd.Flags().String("description", "", "A brief description of the fleet to create.")
	iotfleetwise_createFleetCmd.Flags().String("fleet-id", "", "The unique ID of the fleet to create.")
	iotfleetwise_createFleetCmd.Flags().String("signal-catalog-arn", "", "The Amazon Resource Name (ARN) of a signal catalog.")
	iotfleetwise_createFleetCmd.Flags().String("tags", "", "Metadata that can be used to manage the fleet.")
	iotfleetwise_createFleetCmd.MarkFlagRequired("fleet-id")
	iotfleetwise_createFleetCmd.MarkFlagRequired("signal-catalog-arn")
	iotfleetwiseCmd.AddCommand(iotfleetwise_createFleetCmd)
}
