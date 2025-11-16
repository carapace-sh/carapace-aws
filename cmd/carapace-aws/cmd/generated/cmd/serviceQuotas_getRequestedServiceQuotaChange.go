package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getRequestedServiceQuotaChangeCmd = &cobra.Command{
	Use:   "get-requested-service-quota-change",
	Short: "Retrieves information about the specified quota increase request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getRequestedServiceQuotaChangeCmd).Standalone()

	serviceQuotas_getRequestedServiceQuotaChangeCmd.Flags().String("request-id", "", "Specifies the ID of the quota increase request.")
	serviceQuotas_getRequestedServiceQuotaChangeCmd.MarkFlagRequired("request-id")
	serviceQuotasCmd.AddCommand(serviceQuotas_getRequestedServiceQuotaChangeCmd)
}
