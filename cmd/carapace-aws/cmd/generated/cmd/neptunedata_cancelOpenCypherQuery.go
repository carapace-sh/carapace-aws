package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelOpenCypherQueryCmd = &cobra.Command{
	Use:   "cancel-open-cypher-query",
	Short: "Cancels a specified openCypher query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelOpenCypherQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_cancelOpenCypherQueryCmd).Standalone()

		neptunedata_cancelOpenCypherQueryCmd.Flags().Bool("no-silent", false, "If set to `TRUE`, causes the cancelation of the openCypher query to happen silently.")
		neptunedata_cancelOpenCypherQueryCmd.Flags().String("query-id", "", "The unique ID of the openCypher query to cancel.")
		neptunedata_cancelOpenCypherQueryCmd.Flags().Bool("silent", false, "If set to `TRUE`, causes the cancelation of the openCypher query to happen silently.")
		neptunedata_cancelOpenCypherQueryCmd.Flag("no-silent").Hidden = true
		neptunedata_cancelOpenCypherQueryCmd.MarkFlagRequired("query-id")
	})
	neptunedataCmd.AddCommand(neptunedata_cancelOpenCypherQueryCmd)
}
