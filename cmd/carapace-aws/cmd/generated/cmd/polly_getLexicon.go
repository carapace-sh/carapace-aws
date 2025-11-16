package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_getLexiconCmd = &cobra.Command{
	Use:   "get-lexicon",
	Short: "Returns the content of the specified pronunciation lexicon stored in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_getLexiconCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(polly_getLexiconCmd).Standalone()

		polly_getLexiconCmd.Flags().String("name", "", "Name of the lexicon.")
		polly_getLexiconCmd.MarkFlagRequired("name")
	})
	pollyCmd.AddCommand(polly_getLexiconCmd)
}
