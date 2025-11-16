package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_defineSuggesterCmd = &cobra.Command{
	Use:   "define-suggester",
	Short: "Configures a suggester for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_defineSuggesterCmd).Standalone()

	cloudsearch_defineSuggesterCmd.Flags().String("domain-name", "", "")
	cloudsearch_defineSuggesterCmd.Flags().String("suggester", "", "")
	cloudsearch_defineSuggesterCmd.MarkFlagRequired("domain-name")
	cloudsearch_defineSuggesterCmd.MarkFlagRequired("suggester")
	cloudsearchCmd.AddCommand(cloudsearch_defineSuggesterCmd)
}
