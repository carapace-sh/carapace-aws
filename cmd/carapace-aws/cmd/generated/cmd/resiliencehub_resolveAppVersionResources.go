package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_resolveAppVersionResourcesCmd = &cobra.Command{
	Use:   "resolve-app-version-resources",
	Short: "Resolves the resources for an application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_resolveAppVersionResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_resolveAppVersionResourcesCmd).Standalone()

		resiliencehub_resolveAppVersionResourcesCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_resolveAppVersionResourcesCmd.Flags().String("app-version", "", "The version of the application.")
		resiliencehub_resolveAppVersionResourcesCmd.MarkFlagRequired("app-arn")
		resiliencehub_resolveAppVersionResourcesCmd.MarkFlagRequired("app-version")
	})
	resiliencehubCmd.AddCommand(resiliencehub_resolveAppVersionResourcesCmd)
}
