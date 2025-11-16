package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_describeRxNormInferenceJobCmd = &cobra.Command{
	Use:   "describe-rx-norm-inference-job",
	Short: "Gets the properties associated with an InferRxNorm job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_describeRxNormInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_describeRxNormInferenceJobCmd).Standalone()

		comprehendmedical_describeRxNormInferenceJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend Medical generated for the job.")
		comprehendmedical_describeRxNormInferenceJobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_describeRxNormInferenceJobCmd)
}
