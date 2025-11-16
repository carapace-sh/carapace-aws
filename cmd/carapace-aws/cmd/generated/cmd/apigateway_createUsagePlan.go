package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createUsagePlanCmd = &cobra.Command{
	Use:   "create-usage-plan",
	Short: "Creates a usage plan with the throttle and quota limits, as well as the associated API stages, specified in the payload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createUsagePlanCmd).Standalone()

	apigateway_createUsagePlanCmd.Flags().String("api-stages", "", "The associated API stages of the usage plan.")
	apigateway_createUsagePlanCmd.Flags().String("description", "", "The description of the usage plan.")
	apigateway_createUsagePlanCmd.Flags().String("name", "", "The name of the usage plan.")
	apigateway_createUsagePlanCmd.Flags().String("quota", "", "The quota of the usage plan.")
	apigateway_createUsagePlanCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigateway_createUsagePlanCmd.Flags().String("throttle", "", "The throttling limits of the usage plan.")
	apigateway_createUsagePlanCmd.MarkFlagRequired("name")
	apigatewayCmd.AddCommand(apigateway_createUsagePlanCmd)
}
