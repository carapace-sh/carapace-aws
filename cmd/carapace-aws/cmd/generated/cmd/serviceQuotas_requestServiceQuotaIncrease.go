package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_requestServiceQuotaIncreaseCmd = &cobra.Command{
	Use:   "request-service-quota-increase",
	Short: "Submits a quota increase request for the specified quota at the account or resource level.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_requestServiceQuotaIncreaseCmd).Standalone()

	serviceQuotas_requestServiceQuotaIncreaseCmd.Flags().String("context-id", "", "Specifies the resource with an Amazon Resource Name (ARN).")
	serviceQuotas_requestServiceQuotaIncreaseCmd.Flags().String("desired-value", "", "Specifies the new, increased value for the quota.")
	serviceQuotas_requestServiceQuotaIncreaseCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_requestServiceQuotaIncreaseCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_requestServiceQuotaIncreaseCmd.Flags().String("support-case-allowed", "", "Specifies if an Amazon Web Services Support case can be opened for the quota increase request.")
	serviceQuotas_requestServiceQuotaIncreaseCmd.MarkFlagRequired("desired-value")
	serviceQuotas_requestServiceQuotaIncreaseCmd.MarkFlagRequired("quota-code")
	serviceQuotas_requestServiceQuotaIncreaseCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_requestServiceQuotaIncreaseCmd)
}
