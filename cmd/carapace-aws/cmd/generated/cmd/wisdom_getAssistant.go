package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getAssistantCmd = &cobra.Command{
	Use:   "get-assistant",
	Short: "Retrieves information about an assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getAssistantCmd).Standalone()

	wisdom_getAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_getAssistantCmd.MarkFlagRequired("assistant-id")
	wisdomCmd.AddCommand(wisdom_getAssistantCmd)
}
