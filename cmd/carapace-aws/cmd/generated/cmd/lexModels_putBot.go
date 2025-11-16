package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_putBotCmd = &cobra.Command{
	Use:   "put-bot",
	Short: "Creates an Amazon Lex conversational bot or replaces an existing bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_putBotCmd).Standalone()

	lexModels_putBotCmd.Flags().String("abort-statement", "", "When Amazon Lex can't understand the user's input in context, it tries to elicit the information a few times.")
	lexModels_putBotCmd.Flags().String("checksum", "", "Identifies a specific revision of the `$LATEST` version.")
	lexModels_putBotCmd.Flags().Bool("child-directed", false, "For each Amazon Lex bot created with the Amazon Lex Model Building Service, you must specify whether your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to the Children's Online Privacy Protection Act (COPPA) by specifying `true` or `false` in the `childDirected` field.")
	lexModels_putBotCmd.Flags().String("clarification-prompt", "", "When Amazon Lex doesn't understand the user's intent, it uses this message to get clarification.")
	lexModels_putBotCmd.Flags().Bool("create-version", false, "When set to `true` a new numbered version of the bot is created.")
	lexModels_putBotCmd.Flags().String("description", "", "A description of the bot.")
	lexModels_putBotCmd.Flags().Bool("detect-sentiment", false, "When set to `true` user utterances are sent to Amazon Comprehend for sentiment analysis.")
	lexModels_putBotCmd.Flags().Bool("enable-model-improvements", false, "Set to `true` to enable access to natural language understanding improvements.")
	lexModels_putBotCmd.Flags().String("idle-session-ttlin-seconds", "", "The maximum time in seconds that Amazon Lex retains the data gathered in a conversation.")
	lexModels_putBotCmd.Flags().String("intents", "", "An array of `Intent` objects.")
	lexModels_putBotCmd.Flags().String("locale", "", "Specifies the target locale for the bot.")
	lexModels_putBotCmd.Flags().String("name", "", "The name of the bot.")
	lexModels_putBotCmd.Flags().String("nlu-intent-confidence-threshold", "", "Determines the threshold where Amazon Lex will insert the `AMAZON.FallbackIntent`, `AMAZON.KendraSearchIntent`, or both when returning alternative intents in a [PostContent](https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html) or [PostText](https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html) response.")
	lexModels_putBotCmd.Flags().Bool("no-child-directed", false, "For each Amazon Lex bot created with the Amazon Lex Model Building Service, you must specify whether your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to the Children's Online Privacy Protection Act (COPPA) by specifying `true` or `false` in the `childDirected` field.")
	lexModels_putBotCmd.Flags().Bool("no-create-version", false, "When set to `true` a new numbered version of the bot is created.")
	lexModels_putBotCmd.Flags().Bool("no-detect-sentiment", false, "When set to `true` user utterances are sent to Amazon Comprehend for sentiment analysis.")
	lexModels_putBotCmd.Flags().Bool("no-enable-model-improvements", false, "Set to `true` to enable access to natural language understanding improvements.")
	lexModels_putBotCmd.Flags().String("process-behavior", "", "If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run.")
	lexModels_putBotCmd.Flags().String("tags", "", "A list of tags to add to the bot.")
	lexModels_putBotCmd.Flags().String("voice-id", "", "The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user.")
	lexModels_putBotCmd.MarkFlagRequired("child-directed")
	lexModels_putBotCmd.MarkFlagRequired("locale")
	lexModels_putBotCmd.MarkFlagRequired("name")
	lexModels_putBotCmd.Flag("no-child-directed").Hidden = true
	lexModels_putBotCmd.MarkFlagRequired("no-child-directed")
	lexModels_putBotCmd.Flag("no-create-version").Hidden = true
	lexModels_putBotCmd.Flag("no-detect-sentiment").Hidden = true
	lexModels_putBotCmd.Flag("no-enable-model-improvements").Hidden = true
	lexModelsCmd.AddCommand(lexModels_putBotCmd)
}
