package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateAipromptCmd = &cobra.Command{
	Use:   "update-aiprompt",
	Short: "Updates an AI Prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateAipromptCmd).Standalone()

	qconnect_updateAipromptCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI Prompt.")
	qconnect_updateAipromptCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_updateAipromptCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_updateAipromptCmd.Flags().String("description", "", "The description of the Amazon Q in Connect AI Prompt.")
	qconnect_updateAipromptCmd.Flags().String("model-id", "", "The identifier of the model used for this AI Prompt.")
	qconnect_updateAipromptCmd.Flags().String("template-configuration", "", "The configuration of the prompt template for this AI Prompt.")
	qconnect_updateAipromptCmd.Flags().String("visibility-status", "", "The visibility status of the Amazon Q in Connect AI prompt.")
	qconnect_updateAipromptCmd.MarkFlagRequired("ai-prompt-id")
	qconnect_updateAipromptCmd.MarkFlagRequired("assistant-id")
	qconnect_updateAipromptCmd.MarkFlagRequired("visibility-status")
	qconnectCmd.AddCommand(qconnect_updateAipromptCmd)
}
