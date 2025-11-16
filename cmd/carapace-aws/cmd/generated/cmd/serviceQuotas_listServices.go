package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Lists the names and codes for the Amazon Web Services services integrated with Service Quotas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listServicesCmd).Standalone()

	serviceQuotas_listServicesCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
	serviceQuotas_listServicesCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
	serviceQuotasCmd.AddCommand(serviceQuotas_listServicesCmd)
}
