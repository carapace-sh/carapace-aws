package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAssistantCmd = &cobra.Command{
	Use:   "delete-assistant",
	Short: "Deletes an assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAssistantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteAssistantCmd).Standalone()

		qconnect_deleteAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_deleteAssistantCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_deleteAssistantCmd)
}
