package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deleteElasticsearchDomainCmd = &cobra.Command{
	Use:   "delete-elasticsearch-domain",
	Short: "Permanently deletes the specified Elasticsearch domain and all of its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deleteElasticsearchDomainCmd).Standalone()

	es_deleteElasticsearchDomainCmd.Flags().String("domain-name", "", "The name of the Elasticsearch domain that you want to permanently delete.")
	es_deleteElasticsearchDomainCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_deleteElasticsearchDomainCmd)
}
