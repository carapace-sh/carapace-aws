package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getPortalServiceProviderMetadataCmd = &cobra.Command{
	Use:   "get-portal-service-provider-metadata",
	Short: "Gets the service provider metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getPortalServiceProviderMetadataCmd).Standalone()

	workspacesWeb_getPortalServiceProviderMetadataCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_getPortalServiceProviderMetadataCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getPortalServiceProviderMetadataCmd)
}
