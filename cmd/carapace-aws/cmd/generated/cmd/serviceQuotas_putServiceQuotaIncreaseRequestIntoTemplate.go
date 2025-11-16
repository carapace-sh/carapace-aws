package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd = &cobra.Command{
	Use:   "put-service-quota-increase-request-into-template",
	Short: "Adds a quota increase request to your quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd).Standalone()

	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.Flags().String("aws-region", "", "Specifies the Amazon Web Services Region to which the template applies.")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.Flags().String("desired-value", "", "Specifies the new, increased value for the quota.")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.MarkFlagRequired("aws-region")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.MarkFlagRequired("desired-value")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.MarkFlagRequired("quota-code")
	serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_putServiceQuotaIncreaseRequestIntoTemplateCmd)
}
