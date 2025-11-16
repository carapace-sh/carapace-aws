package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd = &cobra.Command{
	Use:   "get-service-quota-increase-request-from-template",
	Short: "Retrieves information about the specified quota increase request in your quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd).Standalone()

	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("aws-region", "", "Specifies the Amazon Web Services Region for which you made the request.")
	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("aws-region")
	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("quota-code")
	serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_getServiceQuotaIncreaseRequestFromTemplateCmd)
}
