package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listAudienceModelsCmd = &cobra.Command{
	Use:   "list-audience-models",
	Short: "Returns a list of audience models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listAudienceModelsCmd).Standalone()

	cleanroomsml_listAudienceModelsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listAudienceModelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listAudienceModelsCmd)
}
