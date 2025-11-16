package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_batchDetectSyntaxCmd = &cobra.Command{
	Use:   "batch-detect-syntax",
	Short: "Inspects the text of a batch of documents for the syntax and part of speech of the words in the document and returns information about them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_batchDetectSyntaxCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_batchDetectSyntaxCmd).Standalone()

		comprehend_batchDetectSyntaxCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_batchDetectSyntaxCmd.Flags().String("text-list", "", "A list containing the UTF-8 encoded text of the input documents.")
		comprehend_batchDetectSyntaxCmd.MarkFlagRequired("language-code")
		comprehend_batchDetectSyntaxCmd.MarkFlagRequired("text-list")
	})
	comprehendCmd.AddCommand(comprehend_batchDetectSyntaxCmd)
}
