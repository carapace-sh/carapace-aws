package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_getResourcesCmd = &cobra.Command{
	Use:   "get-resources",
	Short: "Returns all the tagged or previously tagged resources that are located in the specified Amazon Web Services Region for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_getResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourcegroupstaggingapi_getResourcesCmd).Standalone()

		resourcegroupstaggingapi_getResourcesCmd.Flags().String("exclude-compliant-resources", "", "Specifies whether to exclude resources that are compliant with the tag policy.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("include-compliance-details", "", "Specifies whether to include details regarding the compliance with the effective tag policy.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("pagination-token", "", "Specifies a `PaginationToken` response value from a previous request to indicate that you want the next page of results.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("resource-arnlist", "", "Specifies a list of ARNs of resources for which you want to retrieve tag data.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("resource-type-filters", "", "Specifies the resource types that you want included in the response.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("resources-per-page", "", "Specifies the maximum number of results to be returned in each page.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("tag-filters", "", "Specifies a list of TagFilters (keys and values) to restrict the output to only those resources that have tags with the specified keys and, if included, the specified values.")
		resourcegroupstaggingapi_getResourcesCmd.Flags().String("tags-per-page", "", "Amazon Web Services recommends using `ResourcesPerPage` instead of this parameter.")
	})
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_getResourcesCmd)
}
