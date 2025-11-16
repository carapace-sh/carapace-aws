package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listMedicalVocabulariesCmd = &cobra.Command{
	Use:   "list-medical-vocabularies",
	Short: "Provides a list of custom medical vocabularies that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listMedicalVocabulariesCmd).Standalone()

	transcribe_listMedicalVocabulariesCmd.Flags().String("max-results", "", "The maximum number of custom medical vocabularies to return in each page of results.")
	transcribe_listMedicalVocabulariesCmd.Flags().String("name-contains", "", "Returns only the custom medical vocabularies that contain the specified string.")
	transcribe_listMedicalVocabulariesCmd.Flags().String("next-token", "", "If your `ListMedicalVocabularies` request returns more results than can be displayed, `NextToken` is displayed in the response with an associated string.")
	transcribe_listMedicalVocabulariesCmd.Flags().String("state-equals", "", "Returns only custom medical vocabularies with the specified state.")
	transcribeCmd.AddCommand(transcribe_listMedicalVocabulariesCmd)
}
