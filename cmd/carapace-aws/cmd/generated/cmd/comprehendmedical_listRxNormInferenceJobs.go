package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_listRxNormInferenceJobsCmd = &cobra.Command{
	Use:   "list-rx-norm-inference-jobs",
	Short: "Gets a list of InferRxNorm jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_listRxNormInferenceJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_listRxNormInferenceJobsCmd).Standalone()

		comprehendmedical_listRxNormInferenceJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
		comprehendmedical_listRxNormInferenceJobsCmd.Flags().String("max-results", "", "Identifies the next page of results to return.")
		comprehendmedical_listRxNormInferenceJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_listRxNormInferenceJobsCmd)
}
