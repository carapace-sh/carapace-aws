package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deleteElasticsearchServiceRoleCmd = &cobra.Command{
	Use:   "delete-elasticsearch-service-role",
	Short: "Deletes the service-linked role that Elasticsearch Service uses to manage and maintain VPC domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deleteElasticsearchServiceRoleCmd).Standalone()

	esCmd.AddCommand(es_deleteElasticsearchServiceRoleCmd)
}
