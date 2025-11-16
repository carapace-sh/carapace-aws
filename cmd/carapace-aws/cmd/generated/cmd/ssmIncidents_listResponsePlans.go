package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listResponsePlansCmd = &cobra.Command{
	Use:   "list-response-plans",
	Short: "Lists all response plans in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listResponsePlansCmd).Standalone()

	ssmIncidents_listResponsePlansCmd.Flags().String("max-results", "", "The maximum number of response plans per page.")
	ssmIncidents_listResponsePlansCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	ssmIncidentsCmd.AddCommand(ssmIncidents_listResponsePlansCmd)
}
