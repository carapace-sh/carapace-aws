package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getUsageCmd = &cobra.Command{
	Use:   "get-usage",
	Short: "Gets the usage data of a usage plan in a specified time interval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getUsageCmd).Standalone()

	apigateway_getUsageCmd.Flags().String("end-date", "", "The ending date (e.g., 2016-12-31) of the usage data.")
	apigateway_getUsageCmd.Flags().String("key-id", "", "The Id of the API key associated with the resultant usage data.")
	apigateway_getUsageCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getUsageCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigateway_getUsageCmd.Flags().String("start-date", "", "The starting date (e.g., 2016-01-01) of the usage data.")
	apigateway_getUsageCmd.Flags().String("usage-plan-id", "", "The Id of the usage plan associated with the usage data.")
	apigateway_getUsageCmd.MarkFlagRequired("end-date")
	apigateway_getUsageCmd.MarkFlagRequired("start-date")
	apigateway_getUsageCmd.MarkFlagRequired("usage-plan-id")
	apigatewayCmd.AddCommand(apigateway_getUsageCmd)
}
