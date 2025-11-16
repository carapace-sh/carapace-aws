package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainConfigCmd = &cobra.Command{
	Use:   "describe-domain-config",
	Short: "Returns the configuration of an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeDomainConfigCmd).Standalone()

		opensearch_describeDomainConfigCmd.Flags().String("domain-name", "", "Name of the OpenSearch Service domain configuration that you want to describe.")
		opensearch_describeDomainConfigCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_describeDomainConfigCmd)
}
