package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listTypeVersionsCmd = &cobra.Command{
	Use:   "list-type-versions",
	Short: "Returns summary information about the versions of an extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listTypeVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listTypeVersionsCmd).Standalone()

		cloudformation_listTypeVersionsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension for which you want version summary information.")
		cloudformation_listTypeVersionsCmd.Flags().String("deprecated-status", "", "The deprecation status of the extension versions that you want to get summary information about.")
		cloudformation_listTypeVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listTypeVersionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listTypeVersionsCmd.Flags().String("publisher-id", "", "The publisher ID of the extension publisher.")
		cloudformation_listTypeVersionsCmd.Flags().String("type", "", "The kind of the extension.")
		cloudformation_listTypeVersionsCmd.Flags().String("type-name", "", "The name of the extension for which you want version summary information.")
	})
	cloudformationCmd.AddCommand(cloudformation_listTypeVersionsCmd)
}
