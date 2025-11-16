package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_inferIcd10CmCmd = &cobra.Command{
	Use:   "infer-icd10-cm",
	Short: "InferICD10CM detects medical conditions as entities listed in a patient record and links those entities to normalized concept identifiers in the ICD-10-CM knowledge base from the Centers for Disease Control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_inferIcd10CmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_inferIcd10CmCmd).Standalone()

		comprehendmedical_inferIcd10CmCmd.Flags().String("text", "", "The input text used for analysis.")
		comprehendmedical_inferIcd10CmCmd.MarkFlagRequired("text")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_inferIcd10CmCmd)
}
