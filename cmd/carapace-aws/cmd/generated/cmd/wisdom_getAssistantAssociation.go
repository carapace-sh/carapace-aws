package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getAssistantAssociationCmd = &cobra.Command{
	Use:   "get-assistant-association",
	Short: "Retrieves information about an assistant association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getAssistantAssociationCmd).Standalone()

	wisdom_getAssistantAssociationCmd.Flags().String("assistant-association-id", "", "The identifier of the assistant association.")
	wisdom_getAssistantAssociationCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_getAssistantAssociationCmd.MarkFlagRequired("assistant-association-id")
	wisdom_getAssistantAssociationCmd.MarkFlagRequired("assistant-id")
	wisdomCmd.AddCommand(wisdom_getAssistantAssociationCmd)
}
