package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listDiscoveredResourcesCmd = &cobra.Command{
	Use:   "list-discovered-resources",
	Short: "Returns a list of resource resource identifiers for the specified resource types for the resources of that type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listDiscoveredResourcesCmd).Standalone()

	config_listDiscoveredResourcesCmd.Flags().Bool("include-deleted-resources", false, "Specifies whether Config includes deleted resources in the results.")
	config_listDiscoveredResourcesCmd.Flags().String("limit", "", "The maximum number of resource identifiers returned on each page.")
	config_listDiscoveredResourcesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_listDiscoveredResourcesCmd.Flags().Bool("no-include-deleted-resources", false, "Specifies whether Config includes deleted resources in the results.")
	config_listDiscoveredResourcesCmd.Flags().String("resource-ids", "", "The IDs of only those resources that you want Config to list in the response.")
	config_listDiscoveredResourcesCmd.Flags().String("resource-name", "", "The custom name of only those resources that you want Config to list in the response.")
	config_listDiscoveredResourcesCmd.Flags().String("resource-type", "", "The type of resources that you want Config to list in the response.")
	config_listDiscoveredResourcesCmd.Flag("no-include-deleted-resources").Hidden = true
	config_listDiscoveredResourcesCmd.MarkFlagRequired("resource-type")
	configCmd.AddCommand(config_listDiscoveredResourcesCmd)
}
