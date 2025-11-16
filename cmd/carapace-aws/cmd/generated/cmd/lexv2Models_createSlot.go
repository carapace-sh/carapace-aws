package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createSlotCmd = &cobra.Command{
	Use:   "create-slot",
	Short: "Creates a slot in an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createSlotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createSlotCmd).Standalone()

		lexv2Models_createSlotCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the slot.")
		lexv2Models_createSlotCmd.Flags().String("bot-version", "", "The version of the bot associated with the slot.")
		lexv2Models_createSlotCmd.Flags().String("description", "", "A description of the slot.")
		lexv2Models_createSlotCmd.Flags().String("intent-id", "", "The identifier of the intent that contains the slot.")
		lexv2Models_createSlotCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the slot will be used in.")
		lexv2Models_createSlotCmd.Flags().String("multiple-values-setting", "", "Indicates whether the slot returns multiple values in one response.")
		lexv2Models_createSlotCmd.Flags().String("obfuscation-setting", "", "Determines how slot values are used in Amazon CloudWatch logs.")
		lexv2Models_createSlotCmd.Flags().String("slot-name", "", "The name of the slot.")
		lexv2Models_createSlotCmd.Flags().String("slot-type-id", "", "The unique identifier for the slot type associated with this slot.")
		lexv2Models_createSlotCmd.Flags().String("sub-slot-setting", "", "Specifications for the constituent sub slots and the expression for the composite slot.")
		lexv2Models_createSlotCmd.Flags().String("value-elicitation-setting", "", "Specifies prompts that Amazon Lex sends to the user to elicit a response that provides the value for the slot.")
		lexv2Models_createSlotCmd.MarkFlagRequired("bot-id")
		lexv2Models_createSlotCmd.MarkFlagRequired("bot-version")
		lexv2Models_createSlotCmd.MarkFlagRequired("intent-id")
		lexv2Models_createSlotCmd.MarkFlagRequired("locale-id")
		lexv2Models_createSlotCmd.MarkFlagRequired("slot-name")
		lexv2Models_createSlotCmd.MarkFlagRequired("value-elicitation-setting")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createSlotCmd)
}
