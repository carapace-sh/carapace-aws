package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createIntentCmd = &cobra.Command{
	Use:   "create-intent",
	Short: "Creates an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createIntentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createIntentCmd).Standalone()

		lexv2Models_createIntentCmd.Flags().String("bot-id", "", "The identifier of the bot associated with this intent.")
		lexv2Models_createIntentCmd.Flags().String("bot-version", "", "The version of the bot associated with this intent.")
		lexv2Models_createIntentCmd.Flags().String("description", "", "A description of the intent.")
		lexv2Models_createIntentCmd.Flags().String("dialog-code-hook", "", "Specifies that Amazon Lex invokes the alias Lambda function for each user input.")
		lexv2Models_createIntentCmd.Flags().String("fulfillment-code-hook", "", "Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.")
		lexv2Models_createIntentCmd.Flags().String("initial-response-setting", "", "Configuration settings for the response that is sent to the user at the beginning of a conversation, before eliciting slot values.")
		lexv2Models_createIntentCmd.Flags().String("input-contexts", "", "A list of contexts that must be active for this intent to be considered by Amazon Lex.")
		lexv2Models_createIntentCmd.Flags().String("intent-closing-setting", "", "Sets the response that Amazon Lex sends to the user when the intent is closed.")
		lexv2Models_createIntentCmd.Flags().String("intent-confirmation-setting", "", "Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.")
		lexv2Models_createIntentCmd.Flags().String("intent-name", "", "The name of the intent.")
		lexv2Models_createIntentCmd.Flags().String("kendra-configuration", "", "Configuration information required to use the `AMAZON.KendraSearchIntent` intent to connect to an Amazon Kendra index.")
		lexv2Models_createIntentCmd.Flags().String("locale-id", "", "The identifier of the language and locale where this intent is used.")
		lexv2Models_createIntentCmd.Flags().String("output-contexts", "", "A lists of contexts that the intent activates when it is fulfilled.")
		lexv2Models_createIntentCmd.Flags().String("parent-intent-signature", "", "A unique identifier for the built-in intent to base this intent on.")
		lexv2Models_createIntentCmd.Flags().String("q-in-connect-intent-configuration", "", "Qinconnect intent configuration details for the create intent request.")
		lexv2Models_createIntentCmd.Flags().String("qn-aintent-configuration", "", "Specifies the configuration of the built-in `Amazon.QnAIntent`.")
		lexv2Models_createIntentCmd.Flags().String("sample-utterances", "", "An array of strings that a user might say to signal the intent.")
		lexv2Models_createIntentCmd.MarkFlagRequired("bot-id")
		lexv2Models_createIntentCmd.MarkFlagRequired("bot-version")
		lexv2Models_createIntentCmd.MarkFlagRequired("intent-name")
		lexv2Models_createIntentCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createIntentCmd)
}
