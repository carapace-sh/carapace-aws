package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_listAwsdefaultServiceQuotasCmd = &cobra.Command{
	Use:   "list-awsdefault-service-quotas",
	Short: "Lists the default values for the quotas for the specified Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_listAwsdefaultServiceQuotasCmd).Standalone()

	serviceQuotas_listAwsdefaultServiceQuotasCmd.Flags().String("max-results", "", "Specifies the maximum number of results that you want included on each page of the response.")
	serviceQuotas_listAwsdefaultServiceQuotasCmd.Flags().String("next-token", "", "Specifies a value for receiving additional results after you receive a `NextToken` response in a previous request.")
	serviceQuotas_listAwsdefaultServiceQuotasCmd.Flags().String("service-code", "", "Specifies the service identifier.")
	serviceQuotas_listAwsdefaultServiceQuotasCmd.MarkFlagRequired("service-code")
	serviceQuotasCmd.AddCommand(serviceQuotas_listAwsdefaultServiceQuotasCmd)
}
