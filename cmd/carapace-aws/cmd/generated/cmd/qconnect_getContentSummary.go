package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getContentSummaryCmd = &cobra.Command{
	Use:   "get-content-summary",
	Short: "Retrieves summary information about the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getContentSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getContentSummaryCmd).Standalone()

		qconnect_getContentSummaryCmd.Flags().String("content-id", "", "The identifier of the content.")
		qconnect_getContentSummaryCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_getContentSummaryCmd.MarkFlagRequired("content-id")
		qconnect_getContentSummaryCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_getContentSummaryCmd)
}
