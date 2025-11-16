package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppVersionResourcesCmd = &cobra.Command{
	Use:   "list-app-version-resources",
	Short: "Lists all the resources in an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppVersionResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAppVersionResourcesCmd).Standalone()

		resiliencehub_listAppVersionResourcesCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_listAppVersionResourcesCmd.Flags().String("app-version", "", "The version of the application.")
		resiliencehub_listAppVersionResourcesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listAppVersionResourcesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAppVersionResourcesCmd.Flags().String("resolution-id", "", "The identifier for a specific resolution.")
		resiliencehub_listAppVersionResourcesCmd.MarkFlagRequired("app-arn")
		resiliencehub_listAppVersionResourcesCmd.MarkFlagRequired("app-version")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAppVersionResourcesCmd)
}
