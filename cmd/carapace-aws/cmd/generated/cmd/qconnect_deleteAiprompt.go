package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAipromptCmd = &cobra.Command{
	Use:   "delete-aiprompt",
	Short: "Deletes an Amazon Q in Connect AI Prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAipromptCmd).Standalone()

	qconnect_deleteAipromptCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI prompt.")
	qconnect_deleteAipromptCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_deleteAipromptCmd.MarkFlagRequired("ai-prompt-id")
	qconnect_deleteAipromptCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_deleteAipromptCmd)
}
