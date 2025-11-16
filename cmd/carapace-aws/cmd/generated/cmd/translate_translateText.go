package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_translateTextCmd = &cobra.Command{
	Use:   "translate-text",
	Short: "Translates input text from the source language to the target language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_translateTextCmd).Standalone()

	translate_translateTextCmd.Flags().String("settings", "", "Settings to configure your translation output.")
	translate_translateTextCmd.Flags().String("source-language-code", "", "The language code for the language of the source text.")
	translate_translateTextCmd.Flags().String("target-language-code", "", "The language code requested for the language of the target text.")
	translate_translateTextCmd.Flags().String("terminology-names", "", "The name of a terminology list file to add to the translation job.")
	translate_translateTextCmd.Flags().String("text", "", "The text to translate.")
	translate_translateTextCmd.MarkFlagRequired("source-language-code")
	translate_translateTextCmd.MarkFlagRequired("target-language-code")
	translate_translateTextCmd.MarkFlagRequired("text")
	translateCmd.AddCommand(translate_translateTextCmd)
}
