package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAiagentCmd = &cobra.Command{
	Use:   "create-aiagent",
	Short: "Creates an Amazon Q in Connect AI Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAiagentCmd).Standalone()

	qconnect_createAiagentCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_createAiagentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createAiagentCmd.Flags().String("configuration", "", "The configuration of the AI Agent.")
	qconnect_createAiagentCmd.Flags().String("description", "", "The description of the AI Agent.")
	qconnect_createAiagentCmd.Flags().String("name", "", "The name of the AI Agent.")
	qconnect_createAiagentCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	qconnect_createAiagentCmd.Flags().String("type", "", "The type of the AI Agent.")
	qconnect_createAiagentCmd.Flags().String("visibility-status", "", "The visibility status of the AI Agent.")
	qconnect_createAiagentCmd.MarkFlagRequired("assistant-id")
	qconnect_createAiagentCmd.MarkFlagRequired("configuration")
	qconnect_createAiagentCmd.MarkFlagRequired("name")
	qconnect_createAiagentCmd.MarkFlagRequired("type")
	qconnect_createAiagentCmd.MarkFlagRequired("visibility-status")
	qconnectCmd.AddCommand(qconnect_createAiagentCmd)
}
