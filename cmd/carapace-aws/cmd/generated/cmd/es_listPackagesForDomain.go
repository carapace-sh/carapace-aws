package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listPackagesForDomainCmd = &cobra.Command{
	Use:   "list-packages-for-domain",
	Short: "Lists all packages associated with the Amazon ES domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listPackagesForDomainCmd).Standalone()

	es_listPackagesForDomainCmd.Flags().String("domain-name", "", "The name of the domain for which you want to list associated packages.")
	es_listPackagesForDomainCmd.Flags().String("max-results", "", "Limits results to a maximum number of packages.")
	es_listPackagesForDomainCmd.Flags().String("next-token", "", "Used for pagination.")
	es_listPackagesForDomainCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_listPackagesForDomainCmd)
}
