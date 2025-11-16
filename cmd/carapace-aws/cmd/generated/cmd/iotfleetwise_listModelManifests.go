package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listModelManifestsCmd = &cobra.Command{
	Use:   "list-model-manifests",
	Short: "Retrieves a list of vehicle models (model manifests).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listModelManifestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listModelManifestsCmd).Standalone()

		iotfleetwise_listModelManifestsCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: model manifest name, Amazon Resource Name (ARN), creation time, and last modification time.")
		iotfleetwise_listModelManifestsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listModelManifestsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
		iotfleetwise_listModelManifestsCmd.Flags().String("signal-catalog-arn", "", "The ARN of a signal catalog.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listModelManifestsCmd)
}
