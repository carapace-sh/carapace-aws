package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listSlotTypesCmd = &cobra.Command{
	Use:   "list-slot-types",
	Short: "Gets a list of slot types that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listSlotTypesCmd).Standalone()

	lexv2Models_listSlotTypesCmd.Flags().String("bot-id", "", "The unique identifier of the bot that contains the slot types.")
	lexv2Models_listSlotTypesCmd.Flags().String("bot-version", "", "The version of the bot that contains the slot type.")
	lexv2Models_listSlotTypesCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the slot types in the response to only those that match the filter specification.")
	lexv2Models_listSlotTypesCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the slot types to list.")
	lexv2Models_listSlotTypesCmd.Flags().String("max-results", "", "The maximum number of slot types to return in each page of results.")
	lexv2Models_listSlotTypesCmd.Flags().String("next-token", "", "If the response from the `ListSlotTypes` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listSlotTypesCmd.Flags().String("sort-by", "", "Determines the sort order for the response from the `ListSlotTypes` operation.")
	lexv2Models_listSlotTypesCmd.MarkFlagRequired("bot-id")
	lexv2Models_listSlotTypesCmd.MarkFlagRequired("bot-version")
	lexv2Models_listSlotTypesCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listSlotTypesCmd)
}
