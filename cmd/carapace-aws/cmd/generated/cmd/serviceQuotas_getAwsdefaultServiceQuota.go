package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getAwsdefaultServiceQuotaCmd = &cobra.Command{
	Use:   "get-awsdefault-service-quota",
	Short: "Retrieves the default value for the specified quota.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getAwsdefaultServiceQuotaCmd).Standalone()

	serviceQuotas_getAwsdefaultServiceQuotaCmd.Flags().String("quota-code", "", "Specifies the quota identifier.")
	serviceQuotas_getAwsdefaultServiceQuotaCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_getAwsdefaultServiceQuotaCmd.MarkFlagRequired("quota-code")
	serviceQuotas_getAwsdefaultServiceQuotaCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_getAwsdefaultServiceQuotaCmd)
}
