package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_associateServiceQuotaTemplateCmd = &cobra.Command{
	Use:   "associate-service-quota-template",
	Short: "Associates your quota request template with your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_associateServiceQuotaTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_associateServiceQuotaTemplateCmd).Standalone()

	})
	serviceQuotasCmd.AddCommand(serviceQuotas_associateServiceQuotaTemplateCmd)
}
