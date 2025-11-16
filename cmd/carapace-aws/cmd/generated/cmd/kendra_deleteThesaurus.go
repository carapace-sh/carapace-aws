package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteThesaurusCmd = &cobra.Command{
	Use:   "delete-thesaurus",
	Short: "Deletes an Amazon Kendra thesaurus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteThesaurusCmd).Standalone()

	kendra_deleteThesaurusCmd.Flags().String("id", "", "The identifier of the thesaurus you want to delete.")
	kendra_deleteThesaurusCmd.Flags().String("index-id", "", "The identifier of the index for the thesaurus.")
	kendra_deleteThesaurusCmd.MarkFlagRequired("id")
	kendra_deleteThesaurusCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_deleteThesaurusCmd)
}
