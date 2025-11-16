package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_deleteLexiconCmd = &cobra.Command{
	Use:   "delete-lexicon",
	Short: "Deletes the specified pronunciation lexicon stored in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_deleteLexiconCmd).Standalone()

	polly_deleteLexiconCmd.Flags().String("name", "", "The name of the lexicon to delete.")
	polly_deleteLexiconCmd.MarkFlagRequired("name")
	pollyCmd.AddCommand(polly_deleteLexiconCmd)
}
