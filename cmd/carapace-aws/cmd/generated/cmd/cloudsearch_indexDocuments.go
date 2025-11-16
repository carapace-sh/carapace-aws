package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_indexDocumentsCmd = &cobra.Command{
	Use:   "index-documents",
	Short: "Tells the search domain to start indexing its documents using the latest indexing options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_indexDocumentsCmd).Standalone()

	cloudsearch_indexDocumentsCmd.Flags().String("domain-name", "", "")
	cloudsearch_indexDocumentsCmd.MarkFlagRequired("domain-name")
	cloudsearchCmd.AddCommand(cloudsearch_indexDocumentsCmd)
}
