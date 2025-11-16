package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getApisCmd = &cobra.Command{
	Use:   "get-apis",
	Short: "Gets a collection of Api resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getApisCmd).Standalone()

	apigatewayv2_getApisCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
	apigatewayv2_getApisCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getApisCmd)
}
