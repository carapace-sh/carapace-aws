package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getAssociationForServiceQuotaTemplateCmd = &cobra.Command{
	Use:   "get-association-for-service-quota-template",
	Short: "Retrieves the status of the association for the quota request template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getAssociationForServiceQuotaTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_getAssociationForServiceQuotaTemplateCmd).Standalone()

	})
	serviceQuotasCmd.AddCommand(serviceQuotas_getAssociationForServiceQuotaTemplateCmd)
}
