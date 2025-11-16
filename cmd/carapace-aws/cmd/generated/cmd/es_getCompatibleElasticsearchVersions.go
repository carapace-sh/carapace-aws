package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_getCompatibleElasticsearchVersionsCmd = &cobra.Command{
	Use:   "get-compatible-elasticsearch-versions",
	Short: "Returns a list of upgrade compatible Elastisearch versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_getCompatibleElasticsearchVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_getCompatibleElasticsearchVersionsCmd).Standalone()

		es_getCompatibleElasticsearchVersionsCmd.Flags().String("domain-name", "", "")
	})
	esCmd.AddCommand(es_getCompatibleElasticsearchVersionsCmd)
}
