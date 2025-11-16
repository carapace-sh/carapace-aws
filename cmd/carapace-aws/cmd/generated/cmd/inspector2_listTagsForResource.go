package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags attached to a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listTagsForResourceCmd).Standalone()

		inspector2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon resource number (ARN) of the resource to list tags of.")
		inspector2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	inspector2Cmd.AddCommand(inspector2_listTagsForResourceCmd)
}
