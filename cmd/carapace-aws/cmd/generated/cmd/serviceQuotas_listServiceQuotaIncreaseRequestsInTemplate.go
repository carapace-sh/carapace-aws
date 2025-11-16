package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd = &cobra.Command{
	Use:   "list-service-quota-increase-requests-in-template",
	Short: "Lists the quota increase requests in the specified quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd).Standalone()

		serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd.Flags().String("aws-region", "", "Specifies the Amazon Web Services Region for which you made the request.")
		serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
		serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
		serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	})
	serviceQuotasCmd.AddCommand(serviceQuotas_listServiceQuotaIncreaseRequestsInTemplateCmd)
}
