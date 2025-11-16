package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateModelManifestCmd = &cobra.Command{
	Use:   "update-model-manifest",
	Short: "Updates a vehicle model (model manifest).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateModelManifestCmd).Standalone()

	iotfleetwise_updateModelManifestCmd.Flags().String("description", "", "A brief description of the vehicle model.")
	iotfleetwise_updateModelManifestCmd.Flags().String("name", "", "The name of the vehicle model to update.")
	iotfleetwise_updateModelManifestCmd.Flags().String("nodes-to-add", "", "A list of `fullyQualifiedName` of nodes, which are a general abstraction of signals, to add to the vehicle model.")
	iotfleetwise_updateModelManifestCmd.Flags().String("nodes-to-remove", "", "A list of `fullyQualifiedName` of nodes, which are a general abstraction of signals, to remove from the vehicle model.")
	iotfleetwise_updateModelManifestCmd.Flags().String("status", "", "The state of the vehicle model.")
	iotfleetwise_updateModelManifestCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateModelManifestCmd)
}
