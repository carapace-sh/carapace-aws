package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteSlotTypeCmd = &cobra.Command{
	Use:   "delete-slot-type",
	Short: "Deletes a slot type from a bot locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteSlotTypeCmd).Standalone()

	lexv2Models_deleteSlotTypeCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the slot type.")
	lexv2Models_deleteSlotTypeCmd.Flags().String("bot-version", "", "The version of the bot associated with the slot type.")
	lexv2Models_deleteSlotTypeCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the slot type will be deleted from.")
	lexv2Models_deleteSlotTypeCmd.Flags().String("skip-resource-in-use-check", "", "By default, the `DeleteSlotType` operations throws a `ResourceInUseException` exception if you try to delete a slot type used by a slot.")
	lexv2Models_deleteSlotTypeCmd.Flags().String("slot-type-id", "", "The identifier of the slot type to delete.")
	lexv2Models_deleteSlotTypeCmd.MarkFlagRequired("bot-id")
	lexv2Models_deleteSlotTypeCmd.MarkFlagRequired("bot-version")
	lexv2Models_deleteSlotTypeCmd.MarkFlagRequired("locale-id")
	lexv2Models_deleteSlotTypeCmd.MarkFlagRequired("slot-type-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteSlotTypeCmd)
}
