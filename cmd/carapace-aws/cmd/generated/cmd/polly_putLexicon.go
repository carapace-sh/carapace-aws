package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_putLexiconCmd = &cobra.Command{
	Use:   "put-lexicon",
	Short: "Stores a pronunciation lexicon in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_putLexiconCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(polly_putLexiconCmd).Standalone()

		polly_putLexiconCmd.Flags().String("content", "", "Content of the PLS lexicon as string data.")
		polly_putLexiconCmd.Flags().String("name", "", "Name of the lexicon.")
		polly_putLexiconCmd.MarkFlagRequired("content")
		polly_putLexiconCmd.MarkFlagRequired("name")
	})
	pollyCmd.AddCommand(polly_putLexiconCmd)
}
