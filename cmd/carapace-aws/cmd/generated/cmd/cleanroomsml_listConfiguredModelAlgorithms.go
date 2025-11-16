package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listConfiguredModelAlgorithmsCmd = &cobra.Command{
	Use:   "list-configured-model-algorithms",
	Short: "Returns a list of configured model algorithms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listConfiguredModelAlgorithmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listConfiguredModelAlgorithmsCmd).Standalone()

		cleanroomsml_listConfiguredModelAlgorithmsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listConfiguredModelAlgorithmsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listConfiguredModelAlgorithmsCmd)
}
