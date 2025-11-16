package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getQueryStateCmd = &cobra.Command{
	Use:   "get-query-state",
	Short: "Returns the state of a query previously submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getQueryStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getQueryStateCmd).Standalone()

		lakeformation_getQueryStateCmd.Flags().String("query-id", "", "The ID of the plan query operation.")
		lakeformation_getQueryStateCmd.MarkFlagRequired("query-id")
	})
	lakeformationCmd.AddCommand(lakeformation_getQueryStateCmd)
}
