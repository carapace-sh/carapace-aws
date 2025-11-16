package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_tagResourceCmd).Standalone()

	ecs_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
	ecs_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	ecs_tagResourceCmd.MarkFlagRequired("resource-arn")
	ecs_tagResourceCmd.MarkFlagRequired("tags")
	ecsCmd.AddCommand(ecs_tagResourceCmd)
}
