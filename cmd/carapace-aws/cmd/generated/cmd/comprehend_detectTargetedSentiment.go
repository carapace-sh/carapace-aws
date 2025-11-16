package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectTargetedSentimentCmd = &cobra.Command{
	Use:   "detect-targeted-sentiment",
	Short: "Inspects the input text and returns a sentiment analysis for each entity identified in the text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectTargetedSentimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_detectTargetedSentimentCmd).Standalone()

		comprehend_detectTargetedSentimentCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_detectTargetedSentimentCmd.Flags().String("text", "", "A UTF-8 text string.")
		comprehend_detectTargetedSentimentCmd.MarkFlagRequired("language-code")
		comprehend_detectTargetedSentimentCmd.MarkFlagRequired("text")
	})
	comprehendCmd.AddCommand(comprehend_detectTargetedSentimentCmd)
}
