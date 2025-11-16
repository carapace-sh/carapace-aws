package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listGeneratedTemplatesCmd = &cobra.Command{
	Use:   "list-generated-templates",
	Short: "Lists your generated templates in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listGeneratedTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listGeneratedTemplatesCmd).Standalone()

		cloudformation_listGeneratedTemplatesCmd.Flags().String("max-results", "", "If the number of available results exceeds this maximum, the response includes a `NextToken` value that you can use for the `NextToken` parameter to get the next set of results.")
		cloudformation_listGeneratedTemplatesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	cloudformationCmd.AddCommand(cloudformation_listGeneratedTemplatesCmd)
}
