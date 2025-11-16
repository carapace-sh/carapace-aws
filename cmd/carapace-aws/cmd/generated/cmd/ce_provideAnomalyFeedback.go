package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_provideAnomalyFeedbackCmd = &cobra.Command{
	Use:   "provide-anomaly-feedback",
	Short: "Modifies the feedback property of a given cost anomaly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_provideAnomalyFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_provideAnomalyFeedbackCmd).Standalone()

		ce_provideAnomalyFeedbackCmd.Flags().String("anomaly-id", "", "A cost anomaly ID.")
		ce_provideAnomalyFeedbackCmd.Flags().String("feedback", "", "Describes whether the cost anomaly was a planned activity or you considered it an anomaly.")
		ce_provideAnomalyFeedbackCmd.MarkFlagRequired("anomaly-id")
		ce_provideAnomalyFeedbackCmd.MarkFlagRequired("feedback")
	})
	ceCmd.AddCommand(ce_provideAnomalyFeedbackCmd)
}
