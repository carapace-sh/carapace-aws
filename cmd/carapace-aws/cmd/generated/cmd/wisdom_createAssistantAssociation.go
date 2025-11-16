package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createAssistantAssociationCmd = &cobra.Command{
	Use:   "create-assistant-association",
	Short: "Creates an association between an Amazon Connect Wisdom assistant and another resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createAssistantAssociationCmd).Standalone()

	wisdom_createAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_createAssistantAssociationCmd.Flags().String("association", "", "The identifier of the associated resource.")
	wisdom_createAssistantAssociationCmd.Flags().String("association-type", "", "The type of association.")
	wisdom_createAssistantAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	wisdom_createAssistantAssociationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	wisdom_createAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	wisdom_createAssistantAssociationCmd.MarkFlagRequired("association")
	wisdom_createAssistantAssociationCmd.MarkFlagRequired("association-type")
	wisdomCmd.AddCommand(wisdom_createAssistantAssociationCmd)
}
