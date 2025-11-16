package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_tagResourceCmd).Standalone()

	iottwinmaker_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iottwinmaker_tagResourceCmd.Flags().String("tags", "", "Metadata to add to this resource.")
	iottwinmaker_tagResourceCmd.MarkFlagRequired("resource-arn")
	iottwinmaker_tagResourceCmd.MarkFlagRequired("tags")
	iottwinmakerCmd.AddCommand(iottwinmaker_tagResourceCmd)
}
