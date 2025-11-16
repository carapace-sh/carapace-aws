package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainHealthCmd = &cobra.Command{
	Use:   "describe-domain-health",
	Short: "Returns information about domain and node health, the standby Availability Zone, number of nodes per Availability Zone, and shard count per node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainHealthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeDomainHealthCmd).Standalone()

		opensearch_describeDomainHealthCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_describeDomainHealthCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_describeDomainHealthCmd)
}
