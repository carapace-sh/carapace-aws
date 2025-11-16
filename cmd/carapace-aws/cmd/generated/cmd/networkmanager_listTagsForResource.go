package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listTagsForResourceCmd).Standalone()

	networkmanager_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkmanager_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	networkmanagerCmd.AddCommand(networkmanager_listTagsForResourceCmd)
}
