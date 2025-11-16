package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listDefaultVocabulariesCmd = &cobra.Command{
	Use:   "list-default-vocabularies",
	Short: "Lists the default vocabularies for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listDefaultVocabulariesCmd).Standalone()

	connect_listDefaultVocabulariesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listDefaultVocabulariesCmd.Flags().String("language-code", "", "The language code of the vocabulary entries.")
	connect_listDefaultVocabulariesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listDefaultVocabulariesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listDefaultVocabulariesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listDefaultVocabulariesCmd)
}
