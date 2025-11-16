package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackResourcesCmd = &cobra.Command{
	Use:   "list-stack-resources",
	Short: "Returns descriptions of all resources of the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackResourcesCmd).Standalone()

		cloudformation_listStackResourcesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackResourcesCmd.Flags().String("stack-name", "", "The name or the unique stack ID that is associated with the stack, which aren't always interchangeable:")
		cloudformation_listStackResourcesCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackResourcesCmd)
}
