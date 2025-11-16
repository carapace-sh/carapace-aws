package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_upgradeElasticsearchDomainCmd = &cobra.Command{
	Use:   "upgrade-elasticsearch-domain",
	Short: "Allows you to either upgrade your domain or perform an Upgrade eligibility check to a compatible Elasticsearch version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_upgradeElasticsearchDomainCmd).Standalone()

	es_upgradeElasticsearchDomainCmd.Flags().String("domain-name", "", "")
	es_upgradeElasticsearchDomainCmd.Flags().Bool("no-perform-check-only", false, "This flag, when set to True, indicates that an Upgrade Eligibility Check needs to be performed.")
	es_upgradeElasticsearchDomainCmd.Flags().Bool("perform-check-only", false, "This flag, when set to True, indicates that an Upgrade Eligibility Check needs to be performed.")
	es_upgradeElasticsearchDomainCmd.Flags().String("target-version", "", "The version of Elasticsearch that you intend to upgrade the domain to.")
	es_upgradeElasticsearchDomainCmd.MarkFlagRequired("domain-name")
	es_upgradeElasticsearchDomainCmd.Flag("no-perform-check-only").Hidden = true
	es_upgradeElasticsearchDomainCmd.MarkFlagRequired("target-version")
	esCmd.AddCommand(es_upgradeElasticsearchDomainCmd)
}
