package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listSlotsCmd = &cobra.Command{
	Use:   "list-slots",
	Short: "Gets a list of slots that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listSlotsCmd).Standalone()

	lexv2Models_listSlotsCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the slot.")
	lexv2Models_listSlotsCmd.Flags().String("bot-version", "", "The version of the bot that contains the slot.")
	lexv2Models_listSlotsCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the slots in the response to only those that match the filter specification.")
	lexv2Models_listSlotsCmd.Flags().String("intent-id", "", "The unique identifier of the intent that contains the slot.")
	lexv2Models_listSlotsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the slots to list.")
	lexv2Models_listSlotsCmd.Flags().String("max-results", "", "The maximum number of slots to return in each page of results.")
	lexv2Models_listSlotsCmd.Flags().String("next-token", "", "If the response from the `ListSlots` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listSlotsCmd.Flags().String("sort-by", "", "Determines the sort order for the response from the `ListSlots` operation.")
	lexv2Models_listSlotsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listSlotsCmd.MarkFlagRequired("bot-version")
	lexv2Models_listSlotsCmd.MarkFlagRequired("intent-id")
	lexv2Models_listSlotsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listSlotsCmd)
}
