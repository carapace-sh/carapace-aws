package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_disassociateServiceQuotaTemplateCmd = &cobra.Command{
	Use:   "disassociate-service-quota-template",
	Short: "Disables your quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_disassociateServiceQuotaTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_disassociateServiceQuotaTemplateCmd).Standalone()

	})
	serviceQuotasCmd.AddCommand(serviceQuotas_disassociateServiceQuotaTemplateCmd)
}
