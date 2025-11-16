package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getAssistantCmd = &cobra.Command{
	Use:   "get-assistant",
	Short: "Retrieves information about an assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getAssistantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getAssistantCmd).Standalone()

		qconnect_getAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_getAssistantCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_getAssistantCmd)
}
