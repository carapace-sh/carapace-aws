package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteAudienceGenerationJobCmd = &cobra.Command{
	Use:   "delete-audience-generation-job",
	Short: "Deletes the specified audience generation job, and removes all data associated with the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteAudienceGenerationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteAudienceGenerationJobCmd).Standalone()

		cleanroomsml_deleteAudienceGenerationJobCmd.Flags().String("audience-generation-job-arn", "", "The Amazon Resource Name (ARN) of the audience generation job that you want to delete.")
		cleanroomsml_deleteAudienceGenerationJobCmd.MarkFlagRequired("audience-generation-job-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteAudienceGenerationJobCmd)
}
