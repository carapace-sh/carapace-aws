package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getUsagePlanCmd = &cobra.Command{
	Use:   "get-usage-plan",
	Short: "Gets a usage plan of a given plan identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getUsagePlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getUsagePlanCmd).Standalone()

		apigateway_getUsagePlanCmd.Flags().String("usage-plan-id", "", "The identifier of the UsagePlan resource to be retrieved.")
		apigateway_getUsagePlanCmd.MarkFlagRequired("usage-plan-id")
	})
	apigatewayCmd.AddCommand(apigateway_getUsagePlanCmd)
}
