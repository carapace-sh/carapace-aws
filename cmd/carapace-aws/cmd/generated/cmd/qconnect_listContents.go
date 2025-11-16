package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listContentsCmd = &cobra.Command{
	Use:   "list-contents",
	Short: "Lists the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listContentsCmd).Standalone()

	qconnect_listContentsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_listContentsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listContentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listContentsCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_listContentsCmd)
}
