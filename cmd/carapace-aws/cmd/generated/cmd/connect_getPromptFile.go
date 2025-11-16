package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getPromptFileCmd = &cobra.Command{
	Use:   "get-prompt-file",
	Short: "Gets the prompt file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getPromptFileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getPromptFileCmd).Standalone()

		connect_getPromptFileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getPromptFileCmd.Flags().String("prompt-id", "", "A unique identifier for the prompt.")
		connect_getPromptFileCmd.MarkFlagRequired("instance-id")
		connect_getPromptFileCmd.MarkFlagRequired("prompt-id")
	})
	connectCmd.AddCommand(connect_getPromptFileCmd)
}
