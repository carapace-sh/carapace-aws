package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_startAudienceGenerationJobCmd = &cobra.Command{
	Use:   "start-audience-generation-job",
	Short: "Information necessary to start the audience generation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_startAudienceGenerationJobCmd).Standalone()

	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("collaboration-id", "", "The identifier of the collaboration that contains the audience generation job.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that is used for this audience generation job.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("description", "", "The description of the audience generation job.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().Bool("include-seed-in-output", false, "Whether the seed audience is included in the audience generation output.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("name", "", "The name of the audience generation job.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().Bool("no-include-seed-in-output", false, "Whether the seed audience is included in the audience generation output.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("seed-audience", "", "The seed audience that is used to generate the audience.")
	cleanroomsml_startAudienceGenerationJobCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
	cleanroomsml_startAudienceGenerationJobCmd.MarkFlagRequired("configured-audience-model-arn")
	cleanroomsml_startAudienceGenerationJobCmd.MarkFlagRequired("name")
	cleanroomsml_startAudienceGenerationJobCmd.Flag("no-include-seed-in-output").Hidden = true
	cleanroomsml_startAudienceGenerationJobCmd.MarkFlagRequired("seed-audience")
	cleanroomsmlCmd.AddCommand(cleanroomsml_startAudienceGenerationJobCmd)
}
