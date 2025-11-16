package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAssistantAssociationCmd = &cobra.Command{
	Use:   "delete-assistant-association",
	Short: "Deletes an assistant association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAssistantAssociationCmd).Standalone()

	qconnect_deleteAssistantAssociationCmd.Flags().String("assistant-association-id", "", "The identifier of the assistant association.")
	qconnect_deleteAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_deleteAssistantAssociationCmd.MarkFlagRequired("assistant-association-id")
	qconnect_deleteAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_deleteAssistantAssociationCmd)
}
