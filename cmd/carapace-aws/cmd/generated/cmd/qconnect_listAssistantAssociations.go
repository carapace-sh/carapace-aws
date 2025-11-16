package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAssistantAssociationsCmd = &cobra.Command{
	Use:   "list-assistant-associations",
	Short: "Lists information about assistant associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAssistantAssociationsCmd).Standalone()

	qconnect_listAssistantAssociationsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAssistantAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAssistantAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAssistantAssociationsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAssistantAssociationsCmd)
}
