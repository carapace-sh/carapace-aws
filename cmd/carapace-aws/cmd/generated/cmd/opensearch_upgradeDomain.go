package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_upgradeDomainCmd = &cobra.Command{
	Use:   "upgrade-domain",
	Short: "Allows you to either upgrade your Amazon OpenSearch Service domain or perform an upgrade eligibility check to a compatible version of OpenSearch or Elasticsearch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_upgradeDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_upgradeDomainCmd).Standalone()

		opensearch_upgradeDomainCmd.Flags().String("advanced-options", "", "Only supports the `override_main_response_version` parameter and not other advanced options.")
		opensearch_upgradeDomainCmd.Flags().String("domain-name", "", "Name of the OpenSearch Service domain that you want to upgrade.")
		opensearch_upgradeDomainCmd.Flags().Bool("no-perform-check-only", false, "When true, indicates that an upgrade eligibility check needs to be performed.")
		opensearch_upgradeDomainCmd.Flags().Bool("perform-check-only", false, "When true, indicates that an upgrade eligibility check needs to be performed.")
		opensearch_upgradeDomainCmd.Flags().String("target-version", "", "OpenSearch or Elasticsearch version to which you want to upgrade, in the format Opensearch\\_X.Y or Elasticsearch\\_X.Y.")
		opensearch_upgradeDomainCmd.MarkFlagRequired("domain-name")
		opensearch_upgradeDomainCmd.Flag("no-perform-check-only").Hidden = true
		opensearch_upgradeDomainCmd.MarkFlagRequired("target-version")
	})
	opensearchCmd.AddCommand(opensearch_upgradeDomainCmd)
}
