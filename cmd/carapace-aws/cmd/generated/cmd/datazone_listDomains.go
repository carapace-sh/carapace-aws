package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Lists Amazon DataZone domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDomainsCmd).Standalone()

	datazone_listDomainsCmd.Flags().String("max-results", "", "The maximum number of domains to return in a single call to `ListDomains`.")
	datazone_listDomainsCmd.Flags().String("next-token", "", "When the number of domains is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of domains, the response includes a pagination token named `NextToken`.")
	datazone_listDomainsCmd.Flags().String("status", "", "The status of the data source.")
	datazoneCmd.AddCommand(datazone_listDomainsCmd)
}
