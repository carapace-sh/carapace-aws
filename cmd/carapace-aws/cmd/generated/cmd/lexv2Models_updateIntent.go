package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateIntentCmd = &cobra.Command{
	Use:   "update-intent",
	Short: "Updates the settings for an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateIntentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateIntentCmd).Standalone()

		lexv2Models_updateIntentCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the intent.")
		lexv2Models_updateIntentCmd.Flags().String("bot-version", "", "The version of the bot that contains the intent.")
		lexv2Models_updateIntentCmd.Flags().String("description", "", "The new description of the intent.")
		lexv2Models_updateIntentCmd.Flags().String("dialog-code-hook", "", "The new Lambda function to use between each turn of the conversation with the bot.")
		lexv2Models_updateIntentCmd.Flags().String("fulfillment-code-hook", "", "The new Lambda function to call when all of the intents required slots are provided and the intent is ready for fulfillment.")
		lexv2Models_updateIntentCmd.Flags().String("initial-response-setting", "", "Configuration settings for a response sent to the user before Amazon Lex starts eliciting slots.")
		lexv2Models_updateIntentCmd.Flags().String("input-contexts", "", "A new list of contexts that must be active in order for Amazon Lex to consider the intent.")
		lexv2Models_updateIntentCmd.Flags().String("intent-closing-setting", "", "The new response that Amazon Lex sends the user when the intent is closed.")
		lexv2Models_updateIntentCmd.Flags().String("intent-confirmation-setting", "", "New prompts that Amazon Lex sends to the user to confirm the completion of an intent.")
		lexv2Models_updateIntentCmd.Flags().String("intent-id", "", "The unique identifier of the intent to update.")
		lexv2Models_updateIntentCmd.Flags().String("intent-name", "", "The new name for the intent.")
		lexv2Models_updateIntentCmd.Flags().String("kendra-configuration", "", "New configuration settings for connecting to an Amazon Kendra index.")
		lexv2Models_updateIntentCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this intent is used.")
		lexv2Models_updateIntentCmd.Flags().String("output-contexts", "", "A new list of contexts that Amazon Lex activates when the intent is fulfilled.")
		lexv2Models_updateIntentCmd.Flags().String("parent-intent-signature", "", "The signature of the new built-in intent to use as the parent of this intent.")
		lexv2Models_updateIntentCmd.Flags().String("q-in-connect-intent-configuration", "", "Qinconnect intent configuration details for the update intent request.")
		lexv2Models_updateIntentCmd.Flags().String("qn-aintent-configuration", "", "Specifies the configuration of the built-in `Amazon.QnAIntent`.")
		lexv2Models_updateIntentCmd.Flags().String("sample-utterances", "", "New utterances used to invoke the intent.")
		lexv2Models_updateIntentCmd.Flags().String("slot-priorities", "", "A new list of slots and their priorities that are contained by the intent.")
		lexv2Models_updateIntentCmd.MarkFlagRequired("bot-id")
		lexv2Models_updateIntentCmd.MarkFlagRequired("bot-version")
		lexv2Models_updateIntentCmd.MarkFlagRequired("intent-id")
		lexv2Models_updateIntentCmd.MarkFlagRequired("intent-name")
		lexv2Models_updateIntentCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateIntentCmd)
}
