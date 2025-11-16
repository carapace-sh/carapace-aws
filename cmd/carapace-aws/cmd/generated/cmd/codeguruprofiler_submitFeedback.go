package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_submitFeedbackCmd = &cobra.Command{
	Use:   "submit-feedback",
	Short: "Sends feedback to CodeGuru Profiler about whether the anomaly detected by the analysis is useful or not.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_submitFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_submitFeedbackCmd).Standalone()

		codeguruprofiler_submitFeedbackCmd.Flags().String("anomaly-instance-id", "", "The universally unique identifier (UUID) of the [`AnomalyInstance`](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AnomalyInstance.html) object that is included in the analysis data.")
		codeguruprofiler_submitFeedbackCmd.Flags().String("comment", "", "Optional feedback about this anomaly.")
		codeguruprofiler_submitFeedbackCmd.Flags().String("profiling-group-name", "", "The name of the profiling group that is associated with the analysis data.")
		codeguruprofiler_submitFeedbackCmd.Flags().String("type", "", "The feedback tpye.")
		codeguruprofiler_submitFeedbackCmd.MarkFlagRequired("anomaly-instance-id")
		codeguruprofiler_submitFeedbackCmd.MarkFlagRequired("profiling-group-name")
		codeguruprofiler_submitFeedbackCmd.MarkFlagRequired("type")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_submitFeedbackCmd)
}
