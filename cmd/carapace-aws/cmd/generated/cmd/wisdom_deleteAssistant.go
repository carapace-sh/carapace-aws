package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteAssistantCmd = &cobra.Command{
	Use:   "delete-assistant",
	Short: "Deletes an assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteAssistantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_deleteAssistantCmd).Standalone()

		wisdom_deleteAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
		wisdom_deleteAssistantCmd.MarkFlagRequired("assistant-id")
	})
	wisdomCmd.AddCommand(wisdom_deleteAssistantCmd)
}
