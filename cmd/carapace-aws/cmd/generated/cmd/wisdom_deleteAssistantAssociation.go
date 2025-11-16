package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteAssistantAssociationCmd = &cobra.Command{
	Use:   "delete-assistant-association",
	Short: "Deletes an assistant association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteAssistantAssociationCmd).Standalone()

	wisdom_deleteAssistantAssociationCmd.Flags().String("assistant-association-id", "", "The identifier of the assistant association.")
	wisdom_deleteAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_deleteAssistantAssociationCmd.MarkFlagRequired("assistant-association-id")
	wisdom_deleteAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	wisdomCmd.AddCommand(wisdom_deleteAssistantAssociationCmd)
}
