package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_inferRxNormCmd = &cobra.Command{
	Use:   "infer-rx-norm",
	Short: "InferRxNorm detects medications as entities listed in a patient record and links to the normalized concept identifiers in the RxNorm database from the National Library of Medicine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_inferRxNormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_inferRxNormCmd).Standalone()

		comprehendmedical_inferRxNormCmd.Flags().String("text", "", "The input text used for analysis.")
		comprehendmedical_inferRxNormCmd.MarkFlagRequired("text")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_inferRxNormCmd)
}
