package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_putIntentCmd = &cobra.Command{
	Use:   "put-intent",
	Short: "Creates an intent or replaces an existing intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_putIntentCmd).Standalone()

	lexModels_putIntentCmd.Flags().String("checksum", "", "Identifies a specific revision of the `$LATEST` version.")
	lexModels_putIntentCmd.Flags().String("conclusion-statement", "", "The statement that you want Amazon Lex to convey to the user after the intent is successfully fulfilled by the Lambda function.")
	lexModels_putIntentCmd.Flags().String("confirmation-prompt", "", "Prompts the user to confirm the intent.")
	lexModels_putIntentCmd.Flags().Bool("create-version", false, "When set to `true` a new numbered version of the intent is created.")
	lexModels_putIntentCmd.Flags().String("description", "", "A description of the intent.")
	lexModels_putIntentCmd.Flags().String("dialog-code-hook", "", "Specifies a Lambda function to invoke for each user input.")
	lexModels_putIntentCmd.Flags().String("follow-up-prompt", "", "Amazon Lex uses this prompt to solicit additional activity after fulfilling an intent.")
	lexModels_putIntentCmd.Flags().String("fulfillment-activity", "", "Required.")
	lexModels_putIntentCmd.Flags().String("input-contexts", "", "An array of `InputContext` objects that lists the contexts that must be active for Amazon Lex to choose the intent in a conversation with the user.")
	lexModels_putIntentCmd.Flags().String("kendra-configuration", "", "Configuration information required to use the `AMAZON.KendraSearchIntent` intent to connect to an Amazon Kendra index.")
	lexModels_putIntentCmd.Flags().String("name", "", "The name of the intent.")
	lexModels_putIntentCmd.Flags().Bool("no-create-version", false, "When set to `true` a new numbered version of the intent is created.")
	lexModels_putIntentCmd.Flags().String("output-contexts", "", "An array of `OutputContext` objects that lists the contexts that the intent activates when the intent is fulfilled.")
	lexModels_putIntentCmd.Flags().String("parent-intent-signature", "", "A unique identifier for the built-in intent to base this intent on.")
	lexModels_putIntentCmd.Flags().String("rejection-statement", "", "When the user answers \"no\" to the question defined in `confirmationPrompt`, Amazon Lex responds with this statement to acknowledge that the intent was canceled.")
	lexModels_putIntentCmd.Flags().String("sample-utterances", "", "An array of utterances (strings) that a user might say to signal the intent.")
	lexModels_putIntentCmd.Flags().String("slots", "", "An array of intent slots.")
	lexModels_putIntentCmd.MarkFlagRequired("name")
	lexModels_putIntentCmd.Flag("no-create-version").Hidden = true
	lexModelsCmd.AddCommand(lexModels_putIntentCmd)
}
