package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listDomainNamesCmd = &cobra.Command{
	Use:   "list-domain-names",
	Short: "Returns the names of all Amazon OpenSearch Service domains owned by the current user in the active Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listDomainNamesCmd).Standalone()

	opensearch_listDomainNamesCmd.Flags().String("engine-type", "", "Filters the output by domain engine type.")
	opensearchCmd.AddCommand(opensearch_listDomainNamesCmd)
}
