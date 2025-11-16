package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getModelManifestCmd = &cobra.Command{
	Use:   "get-model-manifest",
	Short: "Retrieves information about a vehicle model (model manifest).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getModelManifestCmd).Standalone()

	iotfleetwise_getModelManifestCmd.Flags().String("name", "", "The name of the vehicle model to retrieve information about.")
	iotfleetwise_getModelManifestCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_getModelManifestCmd)
}
