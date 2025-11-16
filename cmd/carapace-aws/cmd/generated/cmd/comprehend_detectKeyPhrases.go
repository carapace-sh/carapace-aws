package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectKeyPhrasesCmd = &cobra.Command{
	Use:   "detect-key-phrases",
	Short: "Detects the key noun phrases found in the text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectKeyPhrasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_detectKeyPhrasesCmd).Standalone()

		comprehend_detectKeyPhrasesCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_detectKeyPhrasesCmd.Flags().String("text", "", "A UTF-8 text string.")
		comprehend_detectKeyPhrasesCmd.MarkFlagRequired("language-code")
		comprehend_detectKeyPhrasesCmd.MarkFlagRequired("text")
	})
	comprehendCmd.AddCommand(comprehend_detectKeyPhrasesCmd)
}
