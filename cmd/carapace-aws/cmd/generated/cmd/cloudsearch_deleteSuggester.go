package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_deleteSuggesterCmd = &cobra.Command{
	Use:   "delete-suggester",
	Short: "Deletes a suggester.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_deleteSuggesterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_deleteSuggesterCmd).Standalone()

		cloudsearch_deleteSuggesterCmd.Flags().String("domain-name", "", "")
		cloudsearch_deleteSuggesterCmd.Flags().String("suggester-name", "", "Specifies the name of the suggester you want to delete.")
		cloudsearch_deleteSuggesterCmd.MarkFlagRequired("domain-name")
		cloudsearch_deleteSuggesterCmd.MarkFlagRequired("suggester-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_deleteSuggesterCmd)
}
