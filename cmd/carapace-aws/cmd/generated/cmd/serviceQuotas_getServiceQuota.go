package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getServiceQuotaCmd = &cobra.Command{
	Use:   "get-service-quota",
	Short: "Retrieves the applied quota value for the specified account-level or resource-level quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getServiceQuotaCmd).Standalone()

	serviceQuotas_getServiceQuotaCmd.Flags().String("context-id", "", "Specifies the resource with an Amazon Resource Name (ARN).")
	serviceQuotas_getServiceQuotaCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_getServiceQuotaCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_getServiceQuotaCmd.MarkFlagRequired("quota-code")
	serviceQuotas_getServiceQuotaCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_getServiceQuotaCmd)
}
