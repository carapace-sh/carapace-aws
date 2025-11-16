package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createBotVersionCmd = &cobra.Command{
	Use:   "create-bot-version",
	Short: "Creates an immutable version of the bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createBotVersionCmd).Standalone()

	lexv2Models_createBotVersionCmd.Flags().String("bot-id", "", "The identifier of the bot to create the version for.")
	lexv2Models_createBotVersionCmd.Flags().String("bot-version-locale-specification", "", "Specifies the locales that Amazon Lex adds to this version.")
	lexv2Models_createBotVersionCmd.Flags().String("description", "", "A description of the version.")
	lexv2Models_createBotVersionCmd.MarkFlagRequired("bot-id")
	lexv2Models_createBotVersionCmd.MarkFlagRequired("bot-version-locale-specification")
	lexv2ModelsCmd.AddCommand(lexv2Models_createBotVersionCmd)
}
