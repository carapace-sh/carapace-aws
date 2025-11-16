package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getIndexingRulesCmd = &cobra.Command{
	Use:   "get-indexing-rules",
	Short: "Retrieves all indexing rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getIndexingRulesCmd).Standalone()

	xray_getIndexingRulesCmd.Flags().String("next-token", "", "Specify the pagination token returned by a previous request to retrieve the next page of indexes.")
	xrayCmd.AddCommand(xray_getIndexingRulesCmd)
}
