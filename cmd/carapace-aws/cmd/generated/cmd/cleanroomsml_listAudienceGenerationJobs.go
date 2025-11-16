package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listAudienceGenerationJobsCmd = &cobra.Command{
	Use:   "list-audience-generation-jobs",
	Short: "Returns a list of audience generation jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listAudienceGenerationJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listAudienceGenerationJobsCmd).Standalone()

		cleanroomsml_listAudienceGenerationJobsCmd.Flags().String("collaboration-id", "", "The identifier of the collaboration that contains the audience generation jobs that you are interested in.")
		cleanroomsml_listAudienceGenerationJobsCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that was used for the audience generation jobs that you are interested in.")
		cleanroomsml_listAudienceGenerationJobsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listAudienceGenerationJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listAudienceGenerationJobsCmd)
}
