package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteModelManifestCmd = &cobra.Command{
	Use:   "delete-model-manifest",
	Short: "Deletes a vehicle model (model manifest).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteModelManifestCmd).Standalone()

	iotfleetwise_deleteModelManifestCmd.Flags().String("name", "", "The name of the model manifest to delete.")
	iotfleetwise_deleteModelManifestCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteModelManifestCmd)
}
