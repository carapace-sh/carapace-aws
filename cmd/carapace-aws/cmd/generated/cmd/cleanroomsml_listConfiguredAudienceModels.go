package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listConfiguredAudienceModelsCmd = &cobra.Command{
	Use:   "list-configured-audience-models",
	Short: "Returns a list of the configured audience models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listConfiguredAudienceModelsCmd).Standalone()

	cleanroomsml_listConfiguredAudienceModelsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listConfiguredAudienceModelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listConfiguredAudienceModelsCmd)
}
