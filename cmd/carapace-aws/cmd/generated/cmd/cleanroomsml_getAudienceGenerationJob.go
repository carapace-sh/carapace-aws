package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getAudienceGenerationJobCmd = &cobra.Command{
	Use:   "get-audience-generation-job",
	Short: "Returns information about an audience generation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getAudienceGenerationJobCmd).Standalone()

	cleanroomsml_getAudienceGenerationJobCmd.Flags().String("audience-generation-job-arn", "", "The Amazon Resource Name (ARN) of the audience generation job that you are interested in.")
	cleanroomsml_getAudienceGenerationJobCmd.MarkFlagRequired("audience-generation-job-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getAudienceGenerationJobCmd)
}
