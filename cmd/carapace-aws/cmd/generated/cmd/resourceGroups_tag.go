package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Adds tags to a resource group with the specified Amazon resource name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_tagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_tagCmd).Standalone()

		resourceGroups_tagCmd.Flags().String("arn", "", "The Amazon resource name (ARN) of the resource group to which to add tags.")
		resourceGroups_tagCmd.Flags().String("tags", "", "The tags to add to the specified resource group.")
		resourceGroups_tagCmd.MarkFlagRequired("arn")
		resourceGroups_tagCmd.MarkFlagRequired("tags")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_tagCmd)
}
