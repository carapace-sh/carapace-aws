package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectDominantLanguageCmd = &cobra.Command{
	Use:   "batch-detect-dominant-language",
	Short: "Determines the dominant language of the input text for a batch of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectDominantLanguageCmd).Standalone()

	comprehend_batchDetectDominantLanguageCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
	comprehend_batchDetectDominantLanguageCmd.MarkFlagRequired("text-list")
	comprehendCmd.AddCommand(comprehend_batchDetectDominantLanguageCmd)
}
