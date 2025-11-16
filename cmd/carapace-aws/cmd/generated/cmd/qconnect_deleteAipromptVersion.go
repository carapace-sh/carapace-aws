package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAipromptVersionCmd = &cobra.Command{
	Use:   "delete-aiprompt-version",
	Short: "Delete and Amazon Q in Connect AI Prompt version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAipromptVersionCmd).Standalone()

	qconnect_deleteAipromptVersionCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI prompt.")
	qconnect_deleteAipromptVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_deleteAipromptVersionCmd.Flags().String("version-number", "", "The version number of the AI Prompt version to be deleted.")
	qconnect_deleteAipromptVersionCmd.MarkFlagRequired("ai-prompt-id")
	qconnect_deleteAipromptVersionCmd.MarkFlagRequired("assistant-id")
	qconnect_deleteAipromptVersionCmd.MarkFlagRequired("version-number")
	qconnectCmd.AddCommand(qconnect_deleteAipromptVersionCmd)
}
