package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateUsagePlanCmd = &cobra.Command{
	Use:   "update-usage-plan",
	Short: "Updates a usage plan of a given plan Id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateUsagePlanCmd).Standalone()

	apigateway_updateUsagePlanCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateUsagePlanCmd.Flags().String("usage-plan-id", "", "The Id of the to-be-updated usage plan.")
	apigateway_updateUsagePlanCmd.MarkFlagRequired("usage-plan-id")
	apigatewayCmd.AddCommand(apigateway_updateUsagePlanCmd)
}
