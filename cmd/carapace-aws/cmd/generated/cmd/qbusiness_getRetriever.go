package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getRetrieverCmd = &cobra.Command{
	Use:   "get-retriever",
	Short: "Gets information about an existing retriever used by an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getRetrieverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getRetrieverCmd).Standalone()

		qbusiness_getRetrieverCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application using the retriever.")
		qbusiness_getRetrieverCmd.Flags().String("retriever-id", "", "The identifier of the retriever.")
		qbusiness_getRetrieverCmd.MarkFlagRequired("application-id")
		qbusiness_getRetrieverCmd.MarkFlagRequired("retriever-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getRetrieverCmd)
}
