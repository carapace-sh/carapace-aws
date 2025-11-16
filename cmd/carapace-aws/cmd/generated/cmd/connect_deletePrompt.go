package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deletePromptCmd = &cobra.Command{
	Use:   "delete-prompt",
	Short: "Deletes a prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deletePromptCmd).Standalone()

	connect_deletePromptCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deletePromptCmd.Flags().String("prompt-id", "", "A unique identifier for the prompt.")
	connect_deletePromptCmd.MarkFlagRequired("instance-id")
	connect_deletePromptCmd.MarkFlagRequired("prompt-id")
	connectCmd.AddCommand(connect_deletePromptCmd)
}
