package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteAccessLogSettingsCmd = &cobra.Command{
	Use:   "delete-access-log-settings",
	Short: "Deletes the AccessLogSettings for a Stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteAccessLogSettingsCmd).Standalone()

	apigatewayv2_deleteAccessLogSettingsCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteAccessLogSettingsCmd.Flags().String("stage-name", "", "The stage name.")
	apigatewayv2_deleteAccessLogSettingsCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteAccessLogSettingsCmd.MarkFlagRequired("stage-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteAccessLogSettingsCmd)
}
