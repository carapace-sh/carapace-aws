package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getCompatibleVersionsCmd = &cobra.Command{
	Use:   "get-compatible-versions",
	Short: "Returns a map of OpenSearch or Elasticsearch versions and the versions you can upgrade them to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getCompatibleVersionsCmd).Standalone()

	opensearch_getCompatibleVersionsCmd.Flags().String("domain-name", "", "The name of an existing domain.")
	opensearchCmd.AddCommand(opensearch_getCompatibleVersionsCmd)
}
