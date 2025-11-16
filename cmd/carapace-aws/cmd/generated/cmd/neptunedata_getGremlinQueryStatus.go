package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getGremlinQueryStatusCmd = &cobra.Command{
	Use:   "get-gremlin-query-status",
	Short: "Gets the status of a specified Gremlin query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getGremlinQueryStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getGremlinQueryStatusCmd).Standalone()

		neptunedata_getGremlinQueryStatusCmd.Flags().String("query-id", "", "The unique identifier that identifies the Gremlin query.")
		neptunedata_getGremlinQueryStatusCmd.MarkFlagRequired("query-id")
	})
	neptunedataCmd.AddCommand(neptunedata_getGremlinQueryStatusCmd)
}
