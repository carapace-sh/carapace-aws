package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listAudienceExportJobsCmd = &cobra.Command{
	Use:   "list-audience-export-jobs",
	Short: "Returns a list of the audience export jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listAudienceExportJobsCmd).Standalone()

	cleanroomsml_listAudienceExportJobsCmd.Flags().String("audience-generation-job-arn", "", "The Amazon Resource Name (ARN) of the audience generation job that you are interested in.")
	cleanroomsml_listAudienceExportJobsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listAudienceExportJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listAudienceExportJobsCmd)
}
