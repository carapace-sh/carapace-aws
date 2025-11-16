package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes an Amazon OpenSearch Service domain and all of its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteDomainCmd).Standalone()

	opensearch_deleteDomainCmd.Flags().String("domain-name", "", "The name of the domain you want to permanently delete.")
	opensearch_deleteDomainCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_deleteDomainCmd)
}
