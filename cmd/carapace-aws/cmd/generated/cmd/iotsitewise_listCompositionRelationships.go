package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listCompositionRelationshipsCmd = &cobra.Command{
	Use:   "list-composition-relationships",
	Short: "Retrieves a paginated list of composition relationships for an asset model of type `COMPONENT_MODEL`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listCompositionRelationshipsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listCompositionRelationshipsCmd).Standalone()

		iotsitewise_listCompositionRelationshipsCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
		iotsitewise_listCompositionRelationshipsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listCompositionRelationshipsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listCompositionRelationshipsCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listCompositionRelationshipsCmd)
}
