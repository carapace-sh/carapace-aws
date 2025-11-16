package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_deleteAppBundleCmd = &cobra.Command{
	Use:   "delete-app-bundle",
	Short: "Deletes an app bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_deleteAppBundleCmd).Standalone()

	appfabric_deleteAppBundleCmd.Flags().String("app-bundle-identifier", "", "The ID or Amazon Resource Name (ARN) of the app bundle that needs to be deleted.")
	appfabric_deleteAppBundleCmd.MarkFlagRequired("app-bundle-identifier")
	appfabricCmd.AddCommand(appfabric_deleteAppBundleCmd)
}
