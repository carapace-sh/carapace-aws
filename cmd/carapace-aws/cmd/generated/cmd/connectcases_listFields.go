package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listFieldsCmd = &cobra.Command{
	Use:   "list-fields",
	Short: "Lists all fields in a Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listFieldsCmd).Standalone()

	connectcases_listFieldsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_listFieldsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listFieldsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listFieldsCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_listFieldsCmd)
}
