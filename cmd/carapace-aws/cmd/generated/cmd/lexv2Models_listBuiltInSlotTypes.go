package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBuiltInSlotTypesCmd = &cobra.Command{
	Use:   "list-built-in-slot-types",
	Short: "Gets a list of built-in slot types that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBuiltInSlotTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBuiltInSlotTypesCmd).Standalone()

		lexv2Models_listBuiltInSlotTypesCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the slot types to list.")
		lexv2Models_listBuiltInSlotTypesCmd.Flags().String("max-results", "", "The maximum number of built-in slot types to return in each page of results.")
		lexv2Models_listBuiltInSlotTypesCmd.Flags().String("next-token", "", "If the response from the `ListBuiltInSlotTypes` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listBuiltInSlotTypesCmd.Flags().String("sort-by", "", "Determines the sort order for the response from the `ListBuiltInSlotTypes` operation.")
		lexv2Models_listBuiltInSlotTypesCmd.MarkFlagRequired("locale-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBuiltInSlotTypesCmd)
}
