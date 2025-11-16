package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_stopRxNormInferenceJobCmd = &cobra.Command{
	Use:   "stop-rx-norm-inference-job",
	Short: "Stops an InferRxNorm inference job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_stopRxNormInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_stopRxNormInferenceJobCmd).Standalone()

		comprehendmedical_stopRxNormInferenceJobCmd.Flags().String("job-id", "", "The identifier of the job.")
		comprehendmedical_stopRxNormInferenceJobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_stopRxNormInferenceJobCmd)
}
