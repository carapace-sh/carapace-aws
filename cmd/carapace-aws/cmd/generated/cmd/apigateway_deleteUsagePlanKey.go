package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteUsagePlanKeyCmd = &cobra.Command{
	Use:   "delete-usage-plan-key",
	Short: "Deletes a usage plan key and remove the underlying API key from the associated usage plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteUsagePlanKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteUsagePlanKeyCmd).Standalone()

		apigateway_deleteUsagePlanKeyCmd.Flags().String("key-id", "", "The Id of the UsagePlanKey resource to be deleted.")
		apigateway_deleteUsagePlanKeyCmd.Flags().String("usage-plan-id", "", "The Id of the UsagePlan resource representing the usage plan containing the to-be-deleted UsagePlanKey resource representing a plan customer.")
		apigateway_deleteUsagePlanKeyCmd.MarkFlagRequired("key-id")
		apigateway_deleteUsagePlanKeyCmd.MarkFlagRequired("usage-plan-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteUsagePlanKeyCmd)
}
