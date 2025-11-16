package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteRetrieverCmd = &cobra.Command{
	Use:   "delete-retriever",
	Short: "Deletes the retriever used by an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteRetrieverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteRetrieverCmd).Standalone()

		qbusiness_deleteRetrieverCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application using the retriever.")
		qbusiness_deleteRetrieverCmd.Flags().String("retriever-id", "", "The identifier of the retriever being deleted.")
		qbusiness_deleteRetrieverCmd.MarkFlagRequired("application-id")
		qbusiness_deleteRetrieverCmd.MarkFlagRequired("retriever-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteRetrieverCmd)
}
