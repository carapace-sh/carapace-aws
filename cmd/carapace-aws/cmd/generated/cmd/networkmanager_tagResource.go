package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_tagResourceCmd).Standalone()

	networkmanager_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkmanager_tagResourceCmd.Flags().String("tags", "", "The tags to apply to the specified resource.")
	networkmanager_tagResourceCmd.MarkFlagRequired("resource-arn")
	networkmanager_tagResourceCmd.MarkFlagRequired("tags")
	networkmanagerCmd.AddCommand(networkmanager_tagResourceCmd)
}
