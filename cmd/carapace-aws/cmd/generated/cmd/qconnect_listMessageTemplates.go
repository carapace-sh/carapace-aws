package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listMessageTemplatesCmd = &cobra.Command{
	Use:   "list-message-templates",
	Short: "Lists all the available Amazon Q in Connect message templates for the specified knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listMessageTemplatesCmd).Standalone()

	qconnect_listMessageTemplatesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_listMessageTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listMessageTemplatesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listMessageTemplatesCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_listMessageTemplatesCmd)
}
