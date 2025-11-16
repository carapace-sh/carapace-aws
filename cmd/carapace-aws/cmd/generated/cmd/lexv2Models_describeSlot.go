package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeSlotCmd = &cobra.Command{
	Use:   "describe-slot",
	Short: "Gets metadata information about a slot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeSlotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeSlotCmd).Standalone()

		lexv2Models_describeSlotCmd.Flags().String("bot-id", "", "The identifier of the bot associated with the slot.")
		lexv2Models_describeSlotCmd.Flags().String("bot-version", "", "The version of the bot associated with the slot.")
		lexv2Models_describeSlotCmd.Flags().String("intent-id", "", "The identifier of the intent that contains the slot.")
		lexv2Models_describeSlotCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the slot to describe.")
		lexv2Models_describeSlotCmd.Flags().String("slot-id", "", "The unique identifier for the slot.")
		lexv2Models_describeSlotCmd.MarkFlagRequired("bot-id")
		lexv2Models_describeSlotCmd.MarkFlagRequired("bot-version")
		lexv2Models_describeSlotCmd.MarkFlagRequired("intent-id")
		lexv2Models_describeSlotCmd.MarkFlagRequired("locale-id")
		lexv2Models_describeSlotCmd.MarkFlagRequired("slot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeSlotCmd)
}
