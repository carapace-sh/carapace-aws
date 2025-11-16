package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createModelManifestCmd = &cobra.Command{
	Use:   "create-model-manifest",
	Short: "Creates a vehicle model (model manifest) that specifies signals (attributes, branches, sensors, and actuators).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createModelManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_createModelManifestCmd).Standalone()

		iotfleetwise_createModelManifestCmd.Flags().String("description", "", "A brief description of the vehicle model.")
		iotfleetwise_createModelManifestCmd.Flags().String("name", "", "The name of the vehicle model to create.")
		iotfleetwise_createModelManifestCmd.Flags().String("nodes", "", "A list of nodes, which are a general abstraction of signals.")
		iotfleetwise_createModelManifestCmd.Flags().String("signal-catalog-arn", "", "The Amazon Resource Name (ARN) of a signal catalog.")
		iotfleetwise_createModelManifestCmd.Flags().String("tags", "", "Metadata that can be used to manage the vehicle model.")
		iotfleetwise_createModelManifestCmd.MarkFlagRequired("name")
		iotfleetwise_createModelManifestCmd.MarkFlagRequired("nodes")
		iotfleetwise_createModelManifestCmd.MarkFlagRequired("signal-catalog-arn")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_createModelManifestCmd)
}
