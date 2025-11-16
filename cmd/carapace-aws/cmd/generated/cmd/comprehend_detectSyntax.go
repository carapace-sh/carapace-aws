package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_detectSyntaxCmd = &cobra.Command{
	Use:   "detect-syntax",
	Short: "Inspects text for syntax and the part of speech of words in the document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_detectSyntaxCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_detectSyntaxCmd).Standalone()

		comprehend_detectSyntaxCmd.Flags().String("language-code", "", "The language code of the input documents.")
		comprehend_detectSyntaxCmd.Flags().String("text", "", "A UTF-8 string.")
		comprehend_detectSyntaxCmd.MarkFlagRequired("language-code")
		comprehend_detectSyntaxCmd.MarkFlagRequired("text")
	})
	comprehendCmd.AddCommand(comprehend_detectSyntaxCmd)
}
