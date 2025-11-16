package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_listAppBundlesCmd = &cobra.Command{
	Use:   "list-app-bundles",
	Short: "Returns a list of app bundles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_listAppBundlesCmd).Standalone()

	appfabric_listAppBundlesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	appfabric_listAppBundlesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	appfabricCmd.AddCommand(appfabric_listAppBundlesCmd)
}
