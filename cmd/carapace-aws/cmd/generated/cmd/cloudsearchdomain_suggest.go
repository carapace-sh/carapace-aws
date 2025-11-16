package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchdomain_suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Retrieves autocomplete suggestions for a partial query string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchdomain_suggestCmd).Standalone()

	cloudsearchdomain_suggestCmd.Flags().String("query", "", "Specifies the string for which you want to get suggestions.")
	cloudsearchdomain_suggestCmd.Flags().String("size", "", "Specifies the maximum number of suggestions to return.")
	cloudsearchdomain_suggestCmd.Flags().String("suggester", "", "Specifies the name of the suggester to use to find suggested matches.")
	cloudsearchdomain_suggestCmd.MarkFlagRequired("query")
	cloudsearchdomain_suggestCmd.MarkFlagRequired("suggester")
	cloudsearchdomainCmd.AddCommand(cloudsearchdomain_suggestCmd)
}
