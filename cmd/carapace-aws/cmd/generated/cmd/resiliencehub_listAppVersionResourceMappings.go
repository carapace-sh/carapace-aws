package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppVersionResourceMappingsCmd = &cobra.Command{
	Use:   "list-app-version-resource-mappings",
	Short: "Lists how the resources in an application version are mapped/sourced from.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppVersionResourceMappingsCmd).Standalone()

	resiliencehub_listAppVersionResourceMappingsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_listAppVersionResourceMappingsCmd.Flags().String("app-version", "", "The version of the application.")
	resiliencehub_listAppVersionResourceMappingsCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listAppVersionResourceMappingsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppVersionResourceMappingsCmd.MarkFlagRequired("app-arn")
	resiliencehub_listAppVersionResourceMappingsCmd.MarkFlagRequired("app-version")
	resiliencehubCmd.AddCommand(resiliencehub_listAppVersionResourceMappingsCmd)
}
