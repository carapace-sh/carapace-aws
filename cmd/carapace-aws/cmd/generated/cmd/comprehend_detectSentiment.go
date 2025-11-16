package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectSentimentCmd = &cobra.Command{
	Use:   "detect-sentiment",
	Short: "Inspects text and returns an inference of the prevailing sentiment (`POSITIVE`, `NEUTRAL`, `MIXED`, or `NEGATIVE`).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectSentimentCmd).Standalone()

	comprehend_detectSentimentCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehend_detectSentimentCmd.Flags().String("text", "", "A UTF-8 text string.")
	comprehend_detectSentimentCmd.MarkFlagRequired("language-code")
	comprehend_detectSentimentCmd.MarkFlagRequired("text")
	comprehendCmd.AddCommand(comprehend_detectSentimentCmd)
}
