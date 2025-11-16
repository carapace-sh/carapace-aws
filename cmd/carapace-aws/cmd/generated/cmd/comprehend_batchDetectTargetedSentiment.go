package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectTargetedSentimentCmd = &cobra.Command{
	Use:   "batch-detect-targeted-sentiment",
	Short: "Inspects a batch of documents and returns a sentiment analysis for each entity identified in the documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectTargetedSentimentCmd).Standalone()

	comprehend_batchDetectTargetedSentimentCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehend_batchDetectTargetedSentimentCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
	comprehend_batchDetectTargetedSentimentCmd.MarkFlagRequired("language-code")
	comprehend_batchDetectTargetedSentimentCmd.MarkFlagRequired("text-list")
	comprehendCmd.AddCommand(comprehend_batchDetectTargetedSentimentCmd)
}
