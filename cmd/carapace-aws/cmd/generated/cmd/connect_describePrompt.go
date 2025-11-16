package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describePromptCmd = &cobra.Command{
	Use:   "describe-prompt",
	Short: "Describes the prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describePromptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describePromptCmd).Standalone()

		connect_describePromptCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describePromptCmd.Flags().String("prompt-id", "", "A unique identifier for the prompt.")
		connect_describePromptCmd.MarkFlagRequired("instance-id")
		connect_describePromptCmd.MarkFlagRequired("prompt-id")
	})
	connectCmd.AddCommand(connect_describePromptCmd)
}
