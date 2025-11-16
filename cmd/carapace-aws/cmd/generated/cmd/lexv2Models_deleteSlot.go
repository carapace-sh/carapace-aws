package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteSlotCmd = &cobra.Command{
	Use:   "delete-slot",
	Short: "Deletes the specified slot from an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteSlotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteSlotCmd).Standalone()

		lexv2Models_deleteSlotCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the slot to delete.")
		lexv2Models_deleteSlotCmd.Flags().String("bot-version", "", "The version of the bot associated with the slot to delete.")
		lexv2Models_deleteSlotCmd.Flags().String("intent-id", "", "The identifier of the intent associated with the slot.")
		lexv2Models_deleteSlotCmd.Flags().String("locale-id", "", "The identifier of the language and locale that the slot will be deleted from.")
		lexv2Models_deleteSlotCmd.Flags().String("slot-id", "", "The identifier of the slot to delete.")
		lexv2Models_deleteSlotCmd.MarkFlagRequired("bot-id")
		lexv2Models_deleteSlotCmd.MarkFlagRequired("bot-version")
		lexv2Models_deleteSlotCmd.MarkFlagRequired("intent-id")
		lexv2Models_deleteSlotCmd.MarkFlagRequired("locale-id")
		lexv2Models_deleteSlotCmd.MarkFlagRequired("slot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteSlotCmd)
}
