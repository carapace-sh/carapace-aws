package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_tagResourceCmd).Standalone()

	organizations_tagResourceCmd.Flags().String("resource-id", "", "The ID of the resource to add a tag to.")
	organizations_tagResourceCmd.Flags().String("tags", "", "A list of tags to add to the specified resource.")
	organizations_tagResourceCmd.MarkFlagRequired("resource-id")
	organizations_tagResourceCmd.MarkFlagRequired("tags")
	organizationsCmd.AddCommand(organizations_tagResourceCmd)
}
