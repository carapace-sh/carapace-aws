package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getAssistantAssociationCmd = &cobra.Command{
	Use:   "get-assistant-association",
	Short: "Retrieves information about an assistant association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getAssistantAssociationCmd).Standalone()

	qconnect_getAssistantAssociationCmd.Flags().String("assistant-association-id", "", "The identifier of the assistant association.")
	qconnect_getAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_getAssistantAssociationCmd.MarkFlagRequired("assistant-association-id")
	qconnect_getAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_getAssistantAssociationCmd)
}
