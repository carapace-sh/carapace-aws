package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createSlotTypeCmd = &cobra.Command{
	Use:   "create-slot-type",
	Short: "Creates a custom slot type\n\nTo create a custom slot type, specify a name for the slot type and a set of enumeration values, the values that a slot of this type can assume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createSlotTypeCmd).Standalone()

	lexv2Models_createSlotTypeCmd.Flags().String("bot-id", "", "The identifier of the bot associated with this slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("bot-version", "", "The identifier of the bot version associated with this slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("composite-slot-type-setting", "", "Specifications for a composite slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("description", "", "A description of the slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("external-source-setting", "", "Sets the type of external information used to create the slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the slot type will be used in.")
	lexv2Models_createSlotTypeCmd.Flags().String("parent-slot-type-signature", "", "The built-in slot type used as a parent of this slot type.")
	lexv2Models_createSlotTypeCmd.Flags().String("slot-type-name", "", "The name for the slot.")
	lexv2Models_createSlotTypeCmd.Flags().String("slot-type-values", "", "A list of `SlotTypeValue` objects that defines the values that the slot type can take.")
	lexv2Models_createSlotTypeCmd.Flags().String("value-selection-setting", "", "Determines the strategy that Amazon Lex uses to select a value from the list of possible values.")
	lexv2Models_createSlotTypeCmd.MarkFlagRequired("bot-id")
	lexv2Models_createSlotTypeCmd.MarkFlagRequired("bot-version")
	lexv2Models_createSlotTypeCmd.MarkFlagRequired("locale-id")
	lexv2Models_createSlotTypeCmd.MarkFlagRequired("slot-type-name")
	lexv2ModelsCmd.AddCommand(lexv2Models_createSlotTypeCmd)
}
