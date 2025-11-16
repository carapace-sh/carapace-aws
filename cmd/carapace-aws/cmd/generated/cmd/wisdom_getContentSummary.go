package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getContentSummaryCmd = &cobra.Command{
	Use:   "get-content-summary",
	Short: "Retrieves summary information about the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getContentSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_getContentSummaryCmd).Standalone()

		wisdom_getContentSummaryCmd.Flags().String("content-id", "", "The identifier of the content.")
		wisdom_getContentSummaryCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_getContentSummaryCmd.MarkFlagRequired("content-id")
		wisdom_getContentSummaryCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_getContentSummaryCmd)
}
