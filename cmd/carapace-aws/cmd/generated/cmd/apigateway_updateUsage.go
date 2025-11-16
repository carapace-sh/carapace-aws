package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateUsageCmd = &cobra.Command{
	Use:   "update-usage",
	Short: "Grants a temporary extension to the remaining quota of a usage plan associated with a specified API key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateUsageCmd).Standalone()

	apigateway_updateUsageCmd.Flags().String("key-id", "", "The identifier of the API key associated with the usage plan in which a temporary extension is granted to the remaining quota.")
	apigateway_updateUsageCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateUsageCmd.Flags().String("usage-plan-id", "", "The Id of the usage plan associated with the usage data.")
	apigateway_updateUsageCmd.MarkFlagRequired("key-id")
	apigateway_updateUsageCmd.MarkFlagRequired("usage-plan-id")
	apigatewayCmd.AddCommand(apigateway_updateUsageCmd)
}
