package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_createAppBundleCmd = &cobra.Command{
	Use:   "create-app-bundle",
	Short: "Creates an app bundle to collect data from an application using AppFabric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_createAppBundleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_createAppBundleCmd).Standalone()

		appfabric_createAppBundleCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appfabric_createAppBundleCmd.Flags().String("customer-managed-key-identifier", "", "The Amazon Resource Name (ARN) of the Key Management Service (KMS) key to use to encrypt the application data.")
		appfabric_createAppBundleCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	})
	appfabricCmd.AddCommand(appfabric_createAppBundleCmd)
}
