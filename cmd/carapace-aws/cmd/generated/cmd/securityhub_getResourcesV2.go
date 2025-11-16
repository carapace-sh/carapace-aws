package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getResourcesV2Cmd = &cobra.Command{
	Use:   "get-resources-v2",
	Short: "Returns a list of resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getResourcesV2Cmd).Standalone()

	securityhub_getResourcesV2Cmd.Flags().String("filters", "", "Filters resources based on a set of criteria.")
	securityhub_getResourcesV2Cmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_getResourcesV2Cmd.Flags().String("next-token", "", "The token required for pagination.")
	securityhub_getResourcesV2Cmd.Flags().String("sort-criteria", "", "The finding attributes used to sort the list of returned findings.")
	securityhubCmd.AddCommand(securityhub_getResourcesV2Cmd)
}
