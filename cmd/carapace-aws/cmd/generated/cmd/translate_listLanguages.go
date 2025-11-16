package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_listLanguagesCmd = &cobra.Command{
	Use:   "list-languages",
	Short: "Provides a list of languages (RFC-5646 codes and names) that Amazon Translate supports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_listLanguagesCmd).Standalone()

	translate_listLanguagesCmd.Flags().String("display-language-code", "", "The language code for the language to use to display the language names in the response.")
	translate_listLanguagesCmd.Flags().String("max-results", "", "The maximum number of results to return in each response.")
	translate_listLanguagesCmd.Flags().String("next-token", "", "Include the NextToken value to fetch the next group of supported languages.")
	translateCmd.AddCommand(translate_listLanguagesCmd)
}
