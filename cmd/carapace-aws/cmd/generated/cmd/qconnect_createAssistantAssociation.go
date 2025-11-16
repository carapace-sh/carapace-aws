package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAssistantAssociationCmd = &cobra.Command{
	Use:   "create-assistant-association",
	Short: "Creates an association between an Amazon Q in Connect assistant and another resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAssistantAssociationCmd).Standalone()

	qconnect_createAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_createAssistantAssociationCmd.Flags().String("association", "", "The identifier of the associated resource.")
	qconnect_createAssistantAssociationCmd.Flags().String("association-type", "", "The type of association.")
	qconnect_createAssistantAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createAssistantAssociationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	qconnect_createAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	qconnect_createAssistantAssociationCmd.MarkFlagRequired("association")
	qconnect_createAssistantAssociationCmd.MarkFlagRequired("association-type")
	qconnectCmd.AddCommand(qconnect_createAssistantAssociationCmd)
}
