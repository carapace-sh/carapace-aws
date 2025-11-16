package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listServiceQuotasCmd = &cobra.Command{
	Use:   "list-service-quotas",
	Short: "Lists the applied quota values for the specified Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listServiceQuotasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_listServiceQuotasCmd).Standalone()

		serviceQuotas_listServiceQuotasCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
		serviceQuotas_listServiceQuotasCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
		serviceQuotas_listServiceQuotasCmd.Flags().String("quota-applied-at-level", "", "Filters the response to return applied quota values for the `ACCOUNT`, `RESOURCE`, or `ALL` levels.")
		serviceQuotas_listServiceQuotasCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
		serviceQuotas_listServiceQuotasCmd.Flags().String("service-code", "", "Specifies the service identifier.")
		serviceQuotas_listServiceQuotasCmd.MarkFlagRequired("service-code")
	})
	serviceQuotasCmd.AddCommand(serviceQuotas_listServiceQuotasCmd)
}
