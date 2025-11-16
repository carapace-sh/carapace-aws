package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listTagsForResourceCmd).Standalone()

	redshiftServerless_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to list tags for.")
	redshiftServerless_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listTagsForResourceCmd)
}
