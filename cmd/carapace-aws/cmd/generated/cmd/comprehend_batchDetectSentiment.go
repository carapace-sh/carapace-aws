package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectSentimentCmd = &cobra.Command{
	Use:   "batch-detect-sentiment",
	Short: "Inspects a batch of documents and returns an inference of the prevailing sentiment, `POSITIVE`, `NEUTRAL`, `MIXED`, or `NEGATIVE`, in each one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectSentimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_batchDetectSentimentCmd).Standalone()

		comprehend_batchDetectSentimentCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_batchDetectSentimentCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
		comprehend_batchDetectSentimentCmd.MarkFlagRequired("language-code")
		comprehend_batchDetectSentimentCmd.MarkFlagRequired("text-list")
	})
	comprehendCmd.AddCommand(comprehend_batchDetectSentimentCmd)
}
