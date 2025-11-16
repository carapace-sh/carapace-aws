package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_tagResourceCmd).Standalone()

	codeconnections_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to add or update tags.")
	codeconnections_tagResourceCmd.Flags().String("tags", "", "The tags you want to modify or add to the resource.")
	codeconnections_tagResourceCmd.MarkFlagRequired("resource-arn")
	codeconnections_tagResourceCmd.MarkFlagRequired("tags")
	codeconnectionsCmd.AddCommand(codeconnections_tagResourceCmd)
}
