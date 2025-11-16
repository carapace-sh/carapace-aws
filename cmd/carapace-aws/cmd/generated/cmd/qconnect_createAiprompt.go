package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAipromptCmd = &cobra.Command{
	Use:   "create-aiprompt",
	Short: "Creates an Amazon Q in Connect AI Prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAipromptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createAipromptCmd).Standalone()

		qconnect_createAipromptCmd.Flags().String("api-format", "", "The API Format of the AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_createAipromptCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createAipromptCmd.Flags().String("description", "", "The description of the AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("model-id", "", "The identifier of the model used for this AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("name", "", "The name of the AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createAipromptCmd.Flags().String("template-configuration", "", "The configuration of the prompt template for this AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("template-type", "", "The type of the prompt template for this AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("type", "", "The type of this AI Prompt.")
		qconnect_createAipromptCmd.Flags().String("visibility-status", "", "The visibility status of the AI Prompt.")
		qconnect_createAipromptCmd.MarkFlagRequired("api-format")
		qconnect_createAipromptCmd.MarkFlagRequired("assistant-id")
		qconnect_createAipromptCmd.MarkFlagRequired("model-id")
		qconnect_createAipromptCmd.MarkFlagRequired("name")
		qconnect_createAipromptCmd.MarkFlagRequired("template-configuration")
		qconnect_createAipromptCmd.MarkFlagRequired("template-type")
		qconnect_createAipromptCmd.MarkFlagRequired("type")
		qconnect_createAipromptCmd.MarkFlagRequired("visibility-status")
	})
	qconnectCmd.AddCommand(qconnect_createAipromptCmd)
}
