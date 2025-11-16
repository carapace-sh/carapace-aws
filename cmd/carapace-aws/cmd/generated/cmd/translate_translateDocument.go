package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_translateDocumentCmd = &cobra.Command{
	Use:   "translate-document",
	Short: "Translates the input document from the source language to the target language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_translateDocumentCmd).Standalone()

	translate_translateDocumentCmd.Flags().String("document", "", "The content and content type for the document to be translated.")
	translate_translateDocumentCmd.Flags().String("settings", "", "Settings to configure your translation output.")
	translate_translateDocumentCmd.Flags().String("source-language-code", "", "The language code for the language of the source text.")
	translate_translateDocumentCmd.Flags().String("target-language-code", "", "The language code requested for the translated document.")
	translate_translateDocumentCmd.Flags().String("terminology-names", "", "The name of a terminology list file to add to the translation job.")
	translate_translateDocumentCmd.MarkFlagRequired("document")
	translate_translateDocumentCmd.MarkFlagRequired("source-language-code")
	translate_translateDocumentCmd.MarkFlagRequired("target-language-code")
	translateCmd.AddCommand(translate_translateDocumentCmd)
}
