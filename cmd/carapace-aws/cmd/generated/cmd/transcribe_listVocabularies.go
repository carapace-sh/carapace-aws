package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listVocabulariesCmd = &cobra.Command{
	Use:   "list-vocabularies",
	Short: "Provides a list of custom vocabularies that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listVocabulariesCmd).Standalone()

	transcribe_listVocabulariesCmd.Flags().String("max-results", "", "The maximum number of custom vocabularies to return in each page of results.")
	transcribe_listVocabulariesCmd.Flags().String("name-contains", "", "Returns only the custom vocabularies that contain the specified string.")
	transcribe_listVocabulariesCmd.Flags().String("next-token", "", "If your `ListVocabularies` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
	transcribe_listVocabulariesCmd.Flags().String("state-equals", "", "Returns only custom vocabularies with the specified state.")
	transcribeCmd.AddCommand(transcribe_listVocabulariesCmd)
}
