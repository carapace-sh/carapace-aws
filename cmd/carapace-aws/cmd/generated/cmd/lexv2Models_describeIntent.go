package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeIntentCmd = &cobra.Command{
	Use:   "describe-intent",
	Short: "Returns metadata about an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeIntentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeIntentCmd).Standalone()

		lexv2Models_describeIntentCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the intent.")
		lexv2Models_describeIntentCmd.Flags().String("bot-version", "", "The version of the bot associated with the intent.")
		lexv2Models_describeIntentCmd.Flags().String("intent-id", "", "The identifier of the intent to describe.")
		lexv2Models_describeIntentCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the intent to describe.")
		lexv2Models_describeIntentCmd.MarkFlagRequired("bot-id")
		lexv2Models_describeIntentCmd.MarkFlagRequired("bot-version")
		lexv2Models_describeIntentCmd.MarkFlagRequired("intent-id")
		lexv2Models_describeIntentCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeIntentCmd)
}
