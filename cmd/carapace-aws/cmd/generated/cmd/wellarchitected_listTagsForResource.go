package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listTagsForResourceCmd).Standalone()

	wellarchitected_listTagsForResourceCmd.Flags().String("workload-arn", "", "")
	wellarchitected_listTagsForResourceCmd.MarkFlagRequired("workload-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_listTagsForResourceCmd)
}
