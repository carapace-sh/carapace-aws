package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainChangeProgressCmd = &cobra.Command{
	Use:   "describe-domain-change-progress",
	Short: "Returns information about the current blue/green deployment happening on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainChangeProgressCmd).Standalone()

	opensearch_describeDomainChangeProgressCmd.Flags().String("change-id", "", "The specific change ID for which you want to get progress information.")
	opensearch_describeDomainChangeProgressCmd.Flags().String("domain-name", "", "The name of the domain to get progress information for.")
	opensearch_describeDomainChangeProgressCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_describeDomainChangeProgressCmd)
}
