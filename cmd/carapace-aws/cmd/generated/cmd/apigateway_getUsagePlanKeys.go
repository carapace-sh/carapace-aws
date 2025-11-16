package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getUsagePlanKeysCmd = &cobra.Command{
	Use:   "get-usage-plan-keys",
	Short: "Gets all the usage plan keys representing the API keys added to a specified usage plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getUsagePlanKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getUsagePlanKeysCmd).Standalone()

		apigateway_getUsagePlanKeysCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getUsagePlanKeysCmd.Flags().String("name-query", "", "A query parameter specifying the name of the to-be-returned usage plan keys.")
		apigateway_getUsagePlanKeysCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getUsagePlanKeysCmd.Flags().String("usage-plan-id", "", "The Id of the UsagePlan resource representing the usage plan containing the to-be-retrieved UsagePlanKey resource representing a plan customer.")
		apigateway_getUsagePlanKeysCmd.MarkFlagRequired("usage-plan-id")
	})
	apigatewayCmd.AddCommand(apigateway_getUsagePlanKeysCmd)
}
