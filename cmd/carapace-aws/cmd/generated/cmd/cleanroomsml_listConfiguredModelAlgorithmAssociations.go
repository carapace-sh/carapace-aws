package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd = &cobra.Command{
	Use:   "list-configured-model-algorithm-associations",
	Short: "Returns a list of configured model algorithm associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd).Standalone()

		cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd.Flags().String("membership-identifier", "", "The membership ID of the member that created the configured model algorithm associations you are interested in.")
		cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listConfiguredModelAlgorithmAssociationsCmd)
}
