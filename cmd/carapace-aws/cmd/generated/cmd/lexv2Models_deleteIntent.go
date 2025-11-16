package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteIntentCmd = &cobra.Command{
	Use:   "delete-intent",
	Short: "Removes the specified intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteIntentCmd).Standalone()

	lexv2Models_deleteIntentCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the intent.")
	lexv2Models_deleteIntentCmd.Flags().String("bot-version", "", "The version of the bot associated with the intent.")
	lexv2Models_deleteIntentCmd.Flags().String("intent-id", "", "The unique identifier of the intent to delete.")
	lexv2Models_deleteIntentCmd.Flags().String("locale-id", "", "The identifier of the language and locale where the bot will be deleted.")
	lexv2Models_deleteIntentCmd.MarkFlagRequired("bot-id")
	lexv2Models_deleteIntentCmd.MarkFlagRequired("bot-version")
	lexv2Models_deleteIntentCmd.MarkFlagRequired("intent-id")
	lexv2Models_deleteIntentCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteIntentCmd)
}
