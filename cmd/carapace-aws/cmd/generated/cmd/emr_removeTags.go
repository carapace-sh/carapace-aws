package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes tags from an Amazon EMR resource, such as a cluster or Amazon EMR Studio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_removeTagsCmd).Standalone()

	emr_removeTagsCmd.Flags().String("resource-id", "", "The Amazon EMR resource identifier from which tags will be removed.")
	emr_removeTagsCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the resource.")
	emr_removeTagsCmd.MarkFlagRequired("resource-id")
	emr_removeTagsCmd.MarkFlagRequired("tag-keys")
	emrCmd.AddCommand(emr_removeTagsCmd)
}
