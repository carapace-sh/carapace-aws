package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getUsagePlanKeyCmd = &cobra.Command{
	Use:   "get-usage-plan-key",
	Short: "Gets a usage plan key of a given key identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getUsagePlanKeyCmd).Standalone()

	apigateway_getUsagePlanKeyCmd.Flags().String("key-id", "", "The key Id of the to-be-retrieved UsagePlanKey resource representing a plan customer.")
	apigateway_getUsagePlanKeyCmd.Flags().String("usage-plan-id", "", "The Id of the UsagePlan resource representing the usage plan containing the to-be-retrieved UsagePlanKey resource representing a plan customer.")
	apigateway_getUsagePlanKeyCmd.MarkFlagRequired("key-id")
	apigateway_getUsagePlanKeyCmd.MarkFlagRequired("usage-plan-id")
	apigatewayCmd.AddCommand(apigateway_getUsagePlanKeyCmd)
}
