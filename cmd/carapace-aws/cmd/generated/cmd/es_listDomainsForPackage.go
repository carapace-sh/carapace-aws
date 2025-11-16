package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listDomainsForPackageCmd = &cobra.Command{
	Use:   "list-domains-for-package",
	Short: "Lists all Amazon ES domains associated with the package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listDomainsForPackageCmd).Standalone()

	es_listDomainsForPackageCmd.Flags().String("max-results", "", "Limits results to a maximum number of domains.")
	es_listDomainsForPackageCmd.Flags().String("next-token", "", "Used for pagination.")
	es_listDomainsForPackageCmd.Flags().String("package-id", "", "The package for which to list domains.")
	es_listDomainsForPackageCmd.MarkFlagRequired("package-id")
	esCmd.AddCommand(es_listDomainsForPackageCmd)
}
