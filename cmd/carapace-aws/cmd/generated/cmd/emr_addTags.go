package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds tags to an Amazon EMR resource, such as a cluster or an Amazon EMR Studio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_addTagsCmd).Standalone()

	emr_addTagsCmd.Flags().String("resource-id", "", "The Amazon EMR resource identifier to which tags will be added.")
	emr_addTagsCmd.Flags().String("tags", "", "A list of tags to associate with a resource.")
	emr_addTagsCmd.MarkFlagRequired("resource-id")
	emr_addTagsCmd.MarkFlagRequired("tags")
	emrCmd.AddCommand(emr_addTagsCmd)
}
