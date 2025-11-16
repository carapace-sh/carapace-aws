package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listThesauriCmd = &cobra.Command{
	Use:   "list-thesauri",
	Short: "Lists the thesauri for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listThesauriCmd).Standalone()

	kendra_listThesauriCmd.Flags().String("index-id", "", "The identifier of the index with one or more thesauri.")
	kendra_listThesauriCmd.Flags().String("max-results", "", "The maximum number of thesauri to return.")
	kendra_listThesauriCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listThesauriCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listThesauriCmd)
}
