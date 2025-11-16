package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateSlotTypeCmd = &cobra.Command{
	Use:   "update-slot-type",
	Short: "Updates the configuration of an existing slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateSlotTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_updateSlotTypeCmd).Standalone()

		lexv2Models_updateSlotTypeCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("bot-version", "", "The version of the bot that contains the slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("composite-slot-type-setting", "", "Specifications for a composite slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("description", "", "The new description of the slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("external-source-setting", "", "")
		lexv2Models_updateSlotTypeCmd.Flags().String("locale-id", "", "The identifier of the language and locale that contains the slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("parent-slot-type-signature", "", "The new built-in slot type that should be used as the parent of this slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("slot-type-id", "", "The unique identifier of the slot type to update.")
		lexv2Models_updateSlotTypeCmd.Flags().String("slot-type-name", "", "The new name of the slot type.")
		lexv2Models_updateSlotTypeCmd.Flags().String("slot-type-values", "", "A new list of values and their optional synonyms that define the values that the slot type can take.")
		lexv2Models_updateSlotTypeCmd.Flags().String("value-selection-setting", "", "The strategy that Amazon Lex should use when deciding on a value from the list of slot type values.")
		lexv2Models_updateSlotTypeCmd.MarkFlagRequired("bot-id")
		lexv2Models_updateSlotTypeCmd.MarkFlagRequired("bot-version")
		lexv2Models_updateSlotTypeCmd.MarkFlagRequired("locale-id")
		lexv2Models_updateSlotTypeCmd.MarkFlagRequired("slot-type-id")
		lexv2Models_updateSlotTypeCmd.MarkFlagRequired("slot-type-name")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_updateSlotTypeCmd)
}
