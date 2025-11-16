package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves relevant passages or text excerpts given an input query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_retrieveCmd).Standalone()

	kendra_retrieveCmd.Flags().String("attribute-filter", "", "Filters search results by document fields/attributes.")
	kendra_retrieveCmd.Flags().String("document-relevance-override-configurations", "", "Overrides relevance tuning configurations of fields/attributes set at the index level.")
	kendra_retrieveCmd.Flags().String("index-id", "", "The identifier of the index to retrieve relevant passages for the search.")
	kendra_retrieveCmd.Flags().String("page-number", "", "Retrieved relevant passages are returned in pages the size of the `PageSize` parameter.")
	kendra_retrieveCmd.Flags().String("page-size", "", "Sets the number of retrieved relevant passages that are returned in each page of results.")
	kendra_retrieveCmd.Flags().String("query-text", "", "The input query text to retrieve relevant passages for the search.")
	kendra_retrieveCmd.Flags().String("requested-document-attributes", "", "A list of document fields/attributes to include in the response.")
	kendra_retrieveCmd.Flags().String("user-context", "", "The user context token or user and group information.")
	kendra_retrieveCmd.MarkFlagRequired("index-id")
	kendra_retrieveCmd.MarkFlagRequired("query-text")
	kendraCmd.AddCommand(kendra_retrieveCmd)
}
