package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_connectorRegistrationsV2Cmd = &cobra.Command{
	Use:   "connector-registrations-v2",
	Short: "Grants permission to complete the authorization based on input parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_connectorRegistrationsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_connectorRegistrationsV2Cmd).Standalone()

		securityhub_connectorRegistrationsV2Cmd.Flags().String("auth-code", "", "The authCode retrieved from authUrl to complete the OAuth 2.0 authorization code flow.")
		securityhub_connectorRegistrationsV2Cmd.Flags().String("auth-state", "", "The authState retrieved from authUrl to complete the OAuth 2.0 authorization code flow.")
		securityhub_connectorRegistrationsV2Cmd.MarkFlagRequired("auth-code")
		securityhub_connectorRegistrationsV2Cmd.MarkFlagRequired("auth-state")
	})
	securityhubCmd.AddCommand(securityhub_connectorRegistrationsV2Cmd)
}
