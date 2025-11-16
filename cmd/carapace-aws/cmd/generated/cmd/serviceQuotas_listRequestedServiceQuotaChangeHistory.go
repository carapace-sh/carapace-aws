package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd = &cobra.Command{
	Use:   "list-requested-service-quota-change-history",
	Short: "Retrieves the quota increase requests for the specified Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd).Standalone()

		serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
		serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
		serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd.Flags().String("quota-requested-at-level", "", "Filters the response to return quota requests for the `ACCOUNT`, `RESOURCE`, or `ALL` levels.")
		serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd.Flags().String("service-code", "", "Specifies the service identifier.")
		serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd.Flags().String("status", "", "Specifies that you want to filter the results to only the requests with the matching status.")
	})
	serviceQuotasCmd.AddCommand(serviceQuotas_listRequestedServiceQuotaChangeHistoryCmd)
}
