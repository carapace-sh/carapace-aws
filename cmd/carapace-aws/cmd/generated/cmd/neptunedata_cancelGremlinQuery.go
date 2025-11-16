package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelGremlinQueryCmd = &cobra.Command{
	Use:   "cancel-gremlin-query",
	Short: "Cancels a Gremlin query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelGremlinQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_cancelGremlinQueryCmd).Standalone()

		neptunedata_cancelGremlinQueryCmd.Flags().String("query-id", "", "The unique identifier that identifies the query to be canceled.")
		neptunedata_cancelGremlinQueryCmd.MarkFlagRequired("query-id")
	})
	neptunedataCmd.AddCommand(neptunedata_cancelGremlinQueryCmd)
}
