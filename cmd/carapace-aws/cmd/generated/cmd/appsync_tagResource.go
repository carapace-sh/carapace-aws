package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource with user-supplied tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_tagResourceCmd).Standalone()

	appsync_tagResourceCmd.Flags().String("resource-arn", "", "The `GraphqlApi` Amazon Resource Name (ARN).")
	appsync_tagResourceCmd.Flags().String("tags", "", "A `TagMap` object.")
	appsync_tagResourceCmd.MarkFlagRequired("resource-arn")
	appsync_tagResourceCmd.MarkFlagRequired("tags")
	appsyncCmd.AddCommand(appsync_tagResourceCmd)
}
