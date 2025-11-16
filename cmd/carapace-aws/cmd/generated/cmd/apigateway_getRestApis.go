package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getRestApisCmd = &cobra.Command{
	Use:   "get-rest-apis",
	Short: "Lists the RestApis resources for your collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getRestApisCmd).Standalone()

	apigateway_getRestApisCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getRestApisCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigatewayCmd.AddCommand(apigateway_getRestApisCmd)
}
