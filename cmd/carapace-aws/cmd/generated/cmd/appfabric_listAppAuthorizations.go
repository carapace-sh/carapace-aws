package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_listAppAuthorizationsCmd = &cobra.Command{
	Use:   "list-app-authorizations",
	Short: "Returns a list of all app authorizations configured for an app bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_listAppAuthorizationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_listAppAuthorizationsCmd).Standalone()

		appfabric_listAppAuthorizationsCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_listAppAuthorizationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		appfabric_listAppAuthorizationsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		appfabric_listAppAuthorizationsCmd.MarkFlagRequired("app-bundle-identifier")
	})
	appfabricCmd.AddCommand(appfabric_listAppAuthorizationsCmd)
}
