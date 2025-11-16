package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeSlotTypeCmd = &cobra.Command{
	Use:   "describe-slot-type",
	Short: "Gets metadata information about a slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeSlotTypeCmd).Standalone()

	lexv2Models_describeSlotTypeCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the slot type.")
	lexv2Models_describeSlotTypeCmd.Flags().String("bot-version", "", "The version of the bot associated with the slot type.")
	lexv2Models_describeSlotTypeCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the slot type to describe.")
	lexv2Models_describeSlotTypeCmd.Flags().String("slot-type-id", "", "The identifier of the slot type.")
	lexv2Models_describeSlotTypeCmd.MarkFlagRequired("bot-id")
	lexv2Models_describeSlotTypeCmd.MarkFlagRequired("bot-version")
	lexv2Models_describeSlotTypeCmd.MarkFlagRequired("locale-id")
	lexv2Models_describeSlotTypeCmd.MarkFlagRequired("slot-type-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeSlotTypeCmd)
}
