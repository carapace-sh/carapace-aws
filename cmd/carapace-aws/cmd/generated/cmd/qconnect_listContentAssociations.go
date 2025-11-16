package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listContentAssociationsCmd = &cobra.Command{
	Use:   "list-content-associations",
	Short: "Lists the content associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listContentAssociationsCmd).Standalone()

	qconnect_listContentAssociationsCmd.Flags().String("content-id", "", "The identifier of the content.")
	qconnect_listContentAssociationsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_listContentAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listContentAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listContentAssociationsCmd.MarkFlagRequired("content-id")
	qconnect_listContentAssociationsCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_listContentAssociationsCmd)
}
