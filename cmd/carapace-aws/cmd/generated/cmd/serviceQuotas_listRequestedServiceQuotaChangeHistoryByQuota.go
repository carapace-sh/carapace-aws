package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd = &cobra.Command{
	Use:   "list-requested-service-quota-change-history-by-quota",
	Short: "Retrieves the quota increase requests for the specified quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd).Standalone()

	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("quota-requested-at-level", "", "Filters the response to return quota requests for the `ACCOUNT`, `RESOURCE`, or `ALL` levels.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.Flags().String("status", "", "Specifies that you want to filter the results to only the requests with the matching status.")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.MarkFlagRequired("quota-code")
	serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_listRequestedServiceQuotaChangeHistoryByQuotaCmd)
}
