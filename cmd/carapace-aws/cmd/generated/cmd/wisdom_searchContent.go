package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_searchContentCmd = &cobra.Command{
	Use:   "search-content",
	Short: "Searches for content in a specified knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_searchContentCmd).Standalone()

	wisdom_searchContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_searchContentCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_searchContentCmd.Flags().String("next-token", "", "The token for the next set of results.")
	wisdom_searchContentCmd.Flags().String("search-expression", "", "The search expression to filter results.")
	wisdom_searchContentCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_searchContentCmd.MarkFlagRequired("search-expression")
	wisdomCmd.AddCommand(wisdom_searchContentCmd)
}
