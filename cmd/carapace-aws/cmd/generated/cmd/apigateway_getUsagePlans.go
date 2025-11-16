package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getUsagePlansCmd = &cobra.Command{
	Use:   "get-usage-plans",
	Short: "Gets all the usage plans of the caller's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getUsagePlansCmd).Standalone()

	apigateway_getUsagePlansCmd.Flags().String("key-id", "", "The identifier of the API key associated with the usage plans.")
	apigateway_getUsagePlansCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getUsagePlansCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigatewayCmd.AddCommand(apigateway_getUsagePlansCmd)
}
