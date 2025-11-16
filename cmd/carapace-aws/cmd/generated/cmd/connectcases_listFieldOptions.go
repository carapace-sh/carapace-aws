package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listFieldOptionsCmd = &cobra.Command{
	Use:   "list-field-options",
	Short: "Lists all of the field options for a field identifier in the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listFieldOptionsCmd).Standalone()

	connectcases_listFieldOptionsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_listFieldOptionsCmd.Flags().String("field-id", "", "The unique identifier of a field.")
	connectcases_listFieldOptionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listFieldOptionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listFieldOptionsCmd.Flags().String("values", "", "A list of `FieldOption` values to filter on for `ListFieldOptions`.")
	connectcases_listFieldOptionsCmd.MarkFlagRequired("domain-id")
	connectcases_listFieldOptionsCmd.MarkFlagRequired("field-id")
	connectcasesCmd.AddCommand(connectcases_listFieldOptionsCmd)
}
