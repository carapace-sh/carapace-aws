package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_startAudienceExportJobCmd = &cobra.Command{
	Use:   "start-audience-export-job",
	Short: "Export an audience of a specified size after you have generated an audience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_startAudienceExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_startAudienceExportJobCmd).Standalone()

		cleanroomsml_startAudienceExportJobCmd.Flags().String("audience-generation-job-arn", "", "The Amazon Resource Name (ARN) of the audience generation job that you want to export.")
		cleanroomsml_startAudienceExportJobCmd.Flags().String("audience-size", "", "")
		cleanroomsml_startAudienceExportJobCmd.Flags().String("description", "", "The description of the audience export job.")
		cleanroomsml_startAudienceExportJobCmd.Flags().String("name", "", "The name of the audience export job.")
		cleanroomsml_startAudienceExportJobCmd.MarkFlagRequired("audience-generation-job-arn")
		cleanroomsml_startAudienceExportJobCmd.MarkFlagRequired("audience-size")
		cleanroomsml_startAudienceExportJobCmd.MarkFlagRequired("name")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_startAudienceExportJobCmd)
}
