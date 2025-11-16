package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_deleteAppAuthorizationCmd = &cobra.Command{
	Use:   "delete-app-authorization",
	Short: "Deletes an app authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_deleteAppAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_deleteAppAuthorizationCmd).Standalone()

		appfabric_deleteAppAuthorizationCmd.Flags().String("app-authorization-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app authorization to use for the request.")
		appfabric_deleteAppAuthorizationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_deleteAppAuthorizationCmd.MarkFlagRequired("app-authorization-identifier")
		appfabric_deleteAppAuthorizationCmd.MarkFlagRequired("app-bundle-identifier")
	})
	appfabricCmd.AddCommand(appfabric_deleteAppAuthorizationCmd)
}
