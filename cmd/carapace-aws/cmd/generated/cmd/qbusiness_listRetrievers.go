package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listRetrieversCmd = &cobra.Command{
	Use:   "list-retrievers",
	Short: "Lists the retriever used by an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listRetrieversCmd).Standalone()

	qbusiness_listRetrieversCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application using the retriever.")
	qbusiness_listRetrieversCmd.Flags().String("max-results", "", "The maximum number of retrievers returned.")
	qbusiness_listRetrieversCmd.Flags().String("next-token", "", "If the number of retrievers returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of retrievers.")
	qbusiness_listRetrieversCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_listRetrieversCmd)
}
