package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_untagCmd = &cobra.Command{
	Use:   "untag",
	Short: "Deletes tags from a specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_untagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_untagCmd).Standalone()

		resourceGroups_untagCmd.Flags().String("arn", "", "The Amazon resource name (ARN) of the resource group from which to remove tags.")
		resourceGroups_untagCmd.Flags().String("keys", "", "The keys of the tags to be removed.")
		resourceGroups_untagCmd.MarkFlagRequired("arn")
		resourceGroups_untagCmd.MarkFlagRequired("keys")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_untagCmd)
}
