package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_createViewCmd = &cobra.Command{
	Use:   "create-view",
	Short: "Creates a view that users can query by using the [Search]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_createViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_createViewCmd).Standalone()

		resourceExplorer2_createViewCmd.Flags().String("client-token", "", "This value helps ensure idempotency.")
		resourceExplorer2_createViewCmd.Flags().String("filters", "", "An array of strings that specify which resources are included in the results of queries made using this view.")
		resourceExplorer2_createViewCmd.Flags().String("included-properties", "", "Specifies optional fields that you want included in search results from this view.")
		resourceExplorer2_createViewCmd.Flags().String("scope", "", "The root ARN of the account, an organizational unit (OU), or an organization ARN.")
		resourceExplorer2_createViewCmd.Flags().String("tags", "", "Tag key and value pairs that are attached to the view.")
		resourceExplorer2_createViewCmd.Flags().String("view-name", "", "The name of the new view.")
		resourceExplorer2_createViewCmd.MarkFlagRequired("view-name")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_createViewCmd)
}
