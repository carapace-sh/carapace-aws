package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Lists all of the templates in a Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listTemplatesCmd).Standalone()

	connectcases_listTemplatesCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_listTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listTemplatesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listTemplatesCmd.Flags().String("status", "", "A list of status values to filter on.")
	connectcases_listTemplatesCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_listTemplatesCmd)
}
