package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createUsagePlanKeyCmd = &cobra.Command{
	Use:   "create-usage-plan-key",
	Short: "Creates a usage plan key for adding an existing API key to a usage plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createUsagePlanKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createUsagePlanKeyCmd).Standalone()

		apigateway_createUsagePlanKeyCmd.Flags().String("key-id", "", "The identifier of a UsagePlanKey resource for a plan customer.")
		apigateway_createUsagePlanKeyCmd.Flags().String("key-type", "", "The type of a UsagePlanKey resource for a plan customer.")
		apigateway_createUsagePlanKeyCmd.Flags().String("usage-plan-id", "", "The Id of the UsagePlan resource representing the usage plan containing the to-be-created UsagePlanKey resource representing a plan customer.")
		apigateway_createUsagePlanKeyCmd.MarkFlagRequired("key-id")
		apigateway_createUsagePlanKeyCmd.MarkFlagRequired("key-type")
		apigateway_createUsagePlanKeyCmd.MarkFlagRequired("usage-plan-id")
	})
	apigatewayCmd.AddCommand(apigateway_createUsagePlanKeyCmd)
}
