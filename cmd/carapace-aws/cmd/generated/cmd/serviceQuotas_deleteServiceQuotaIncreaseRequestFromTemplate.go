package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd = &cobra.Command{
	Use:   "delete-service-quota-increase-request-from-template",
	Short: "Deletes the quota increase request for the specified quota from your quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd).Standalone()

	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("aws-region", "", "Specifies the Amazon Web Services Region for which the request was made.")
	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("aws-region")
	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("quota-code")
	serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_deleteServiceQuotaIncreaseRequestFromTemplateCmd)
}
