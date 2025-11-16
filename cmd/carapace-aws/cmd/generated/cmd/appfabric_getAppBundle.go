package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_getAppBundleCmd = &cobra.Command{
	Use:   "get-app-bundle",
	Short: "Returns information about an app bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_getAppBundleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_getAppBundleCmd).Standalone()

		appfabric_getAppBundleCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_getAppBundleCmd.MarkFlagRequired("app-bundle-identifier")
	})
	appfabricCmd.AddCommand(appfabric_getAppBundleCmd)
}
