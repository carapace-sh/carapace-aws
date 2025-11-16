package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteCorsConfigurationCmd = &cobra.Command{
	Use:   "delete-cors-configuration",
	Short: "Deletes a CORS configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteCorsConfigurationCmd).Standalone()

	apigatewayv2_deleteCorsConfigurationCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteCorsConfigurationCmd.MarkFlagRequired("api-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteCorsConfigurationCmd)
}
