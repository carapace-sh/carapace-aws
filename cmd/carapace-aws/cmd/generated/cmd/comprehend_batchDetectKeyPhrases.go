package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectKeyPhrasesCmd = &cobra.Command{
	Use:   "batch-detect-key-phrases",
	Short: "Detects the key noun phrases found in a batch of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectKeyPhrasesCmd).Standalone()

	comprehend_batchDetectKeyPhrasesCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehend_batchDetectKeyPhrasesCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
	comprehend_batchDetectKeyPhrasesCmd.MarkFlagRequired("language-code")
	comprehend_batchDetectKeyPhrasesCmd.MarkFlagRequired("text-list")
	comprehendCmd.AddCommand(comprehend_batchDetectKeyPhrasesCmd)
}
