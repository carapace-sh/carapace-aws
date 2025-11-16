package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_untagResourceCmd).Standalone()

	appsync_untagResourceCmd.Flags().String("resource-arn", "", "The `GraphqlApi` Amazon Resource Name (ARN).")
	appsync_untagResourceCmd.Flags().String("tag-keys", "", "A list of `TagKey` objects.")
	appsync_untagResourceCmd.MarkFlagRequired("resource-arn")
	appsync_untagResourceCmd.MarkFlagRequired("tag-keys")
	appsyncCmd.AddCommand(appsync_untagResourceCmd)
}
