package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listTagsForResourceCmd).Standalone()

	connectcases_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN)")
	connectcases_listTagsForResourceCmd.MarkFlagRequired("arn")
	connectcasesCmd.AddCommand(connectcases_listTagsForResourceCmd)
}
