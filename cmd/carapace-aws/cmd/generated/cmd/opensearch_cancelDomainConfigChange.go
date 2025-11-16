package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_cancelDomainConfigChangeCmd = &cobra.Command{
	Use:   "cancel-domain-config-change",
	Short: "Cancels a pending configuration change on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_cancelDomainConfigChangeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_cancelDomainConfigChangeCmd).Standalone()

		opensearch_cancelDomainConfigChangeCmd.Flags().String("domain-name", "", "")
		opensearch_cancelDomainConfigChangeCmd.Flags().String("dry-run", "", "When set to `True`, returns the list of change IDs and properties that will be cancelled without actually cancelling the change.")
		opensearch_cancelDomainConfigChangeCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_cancelDomainConfigChangeCmd)
}
