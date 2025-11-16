package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listSupportedResourceTypesCmd = &cobra.Command{
	Use:   "list-supported-resource-types",
	Short: "Retrieves a list of all resource types currently supported by Amazon Web Services Resource Explorer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listSupportedResourceTypesCmd).Standalone()

	resourceExplorer2_listSupportedResourceTypesCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
	resourceExplorer2_listSupportedResourceTypesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listSupportedResourceTypesCmd)
}
