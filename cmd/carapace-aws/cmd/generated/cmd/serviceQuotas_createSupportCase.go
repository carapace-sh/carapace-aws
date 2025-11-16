package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_createSupportCaseCmd = &cobra.Command{
	Use:   "create-support-case",
	Short: "Creates a Support case for an existing quota increase request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_createSupportCaseCmd).Standalone()

	serviceQuotas_createSupportCaseCmd.Flags().String("request-id", "", "The ID of the pending quota increase request for which you want to open a Support case.")
	serviceQuotas_createSupportCaseCmd.MarkFlagRequired("request-id")
	serviceQuotasCmd.AddCommand(serviceQuotas_createSupportCaseCmd)
}
