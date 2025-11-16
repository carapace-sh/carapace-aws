package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateSlotCmd = &cobra.Command{
	Use:   "update-slot",
	Short: "Updates the settings for a slot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateSlotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateSlotCmd).Standalone()

		lexv2Models_updateSlotCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the slot.")
		lexv2Models_updateSlotCmd.Flags().String("bot-version", "", "The version of the bot that contains the slot.")
		lexv2Models_updateSlotCmd.Flags().String("description", "", "The new description for the slot.")
		lexv2Models_updateSlotCmd.Flags().String("intent-id", "", "The identifier of the intent that contains the slot.")
		lexv2Models_updateSlotCmd.Flags().String("locale-id", "", "The identifier of the language and locale that contains the slot.")
		lexv2Models_updateSlotCmd.Flags().String("multiple-values-setting", "", "Determines whether the slot accepts multiple values in one response.")
		lexv2Models_updateSlotCmd.Flags().String("obfuscation-setting", "", "New settings that determine how slot values are formatted in Amazon CloudWatch logs.")
		lexv2Models_updateSlotCmd.Flags().String("slot-id", "", "The unique identifier for the slot to update.")
		lexv2Models_updateSlotCmd.Flags().String("slot-name", "", "The new name for the slot.")
		lexv2Models_updateSlotCmd.Flags().String("slot-type-id", "", "The unique identifier of the new slot type to associate with this slot.")
		lexv2Models_updateSlotCmd.Flags().String("sub-slot-setting", "", "Specifications for the constituent sub slots and the expression for the composite slot.")
		lexv2Models_updateSlotCmd.Flags().String("value-elicitation-setting", "", "A new set of prompts that Amazon Lex sends to the user to elicit a response the provides a value for the slot.")
		lexv2Models_updateSlotCmd.MarkFlagRequired("bot-id")
		lexv2Models_updateSlotCmd.MarkFlagRequired("bot-version")
		lexv2Models_updateSlotCmd.MarkFlagRequired("intent-id")
		lexv2Models_updateSlotCmd.MarkFlagRequired("locale-id")
		lexv2Models_updateSlotCmd.MarkFlagRequired("slot-id")
		lexv2Models_updateSlotCmd.MarkFlagRequired("slot-name")
		lexv2Models_updateSlotCmd.MarkFlagRequired("value-elicitation-setting")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateSlotCmd)
}
