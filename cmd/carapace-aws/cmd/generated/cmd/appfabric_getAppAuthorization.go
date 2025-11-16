package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_getAppAuthorizationCmd = &cobra.Command{
	Use:   "get-app-authorization",
	Short: "Returns information about an app authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_getAppAuthorizationCmd).Standalone()

	appfabric_getAppAuthorizationCmd.Flags().String("app-authorization-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app authorization to use for the request.")
	appfabric_getAppAuthorizationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_getAppAuthorizationCmd.MarkFlagRequired("app-authorization-identifier")
	appfabric_getAppAuthorizationCmd.MarkFlagRequired("app-bundle-identifier")
	appfabricCmd.AddCommand(appfabric_getAppAuthorizationCmd)
}
