package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listAssistantAssociationsCmd = &cobra.Command{
	Use:   "list-assistant-associations",
	Short: "Lists information about assistant associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listAssistantAssociationsCmd).Standalone()

	wisdom_listAssistantAssociationsCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_listAssistantAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_listAssistantAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	wisdom_listAssistantAssociationsCmd.MarkFlagRequired("assistant-id")
	wisdomCmd.AddCommand(wisdom_listAssistantAssociationsCmd)
}
