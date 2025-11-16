package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_inferSnomedctCmd = &cobra.Command{
	Use:   "infer-snomedct",
	Short: "InferSNOMEDCT detects possible medical concepts as entities and links them to codes from the Systematized Nomenclature of Medicine, Clinical Terms (SNOMED-CT) ontology",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_inferSnomedctCmd).Standalone()

	comprehendmedical_inferSnomedctCmd.Flags().String("text", "", "The input text to be analyzed using InferSNOMEDCT.")
	comprehendmedical_inferSnomedctCmd.MarkFlagRequired("text")
	comprehendmedicalCmd.AddCommand(comprehendmedical_inferSnomedctCmd)
}
