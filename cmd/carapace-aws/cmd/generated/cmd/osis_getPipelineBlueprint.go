package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_getPipelineBlueprintCmd = &cobra.Command{
	Use:   "get-pipeline-blueprint",
	Short: "Retrieves information about a specific blueprint for OpenSearch Ingestion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_getPipelineBlueprintCmd).Standalone()

	osis_getPipelineBlueprintCmd.Flags().String("blueprint-name", "", "The name of the blueprint to retrieve.")
	osis_getPipelineBlueprintCmd.Flags().String("format", "", "The format format of the blueprint to retrieve.")
	osis_getPipelineBlueprintCmd.MarkFlagRequired("blueprint-name")
	osisCmd.AddCommand(osis_getPipelineBlueprintCmd)
}
