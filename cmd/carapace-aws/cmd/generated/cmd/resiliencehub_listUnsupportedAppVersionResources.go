package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listUnsupportedAppVersionResourcesCmd = &cobra.Command{
	Use:   "list-unsupported-app-version-resources",
	Short: "Lists the resources that are not currently supported in Resilience Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listUnsupportedAppVersionResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listUnsupportedAppVersionResourcesCmd).Standalone()

		resiliencehub_listUnsupportedAppVersionResourcesCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.Flags().String("app-version", "", "The version of the application.")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.Flags().String("resolution-id", "", "The identifier for a specific resolution.")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.MarkFlagRequired("app-arn")
		resiliencehub_listUnsupportedAppVersionResourcesCmd.MarkFlagRequired("app-version")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listUnsupportedAppVersionResourcesCmd)
}
