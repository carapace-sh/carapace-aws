package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getAipromptCmd = &cobra.Command{
	Use:   "get-aiprompt",
	Short: "Gets and Amazon Q in Connect AI Prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getAipromptCmd).Standalone()

	qconnect_getAipromptCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI prompt.")
	qconnect_getAipromptCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_getAipromptCmd.MarkFlagRequired("ai-prompt-id")
	qconnect_getAipromptCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_getAipromptCmd)
}
