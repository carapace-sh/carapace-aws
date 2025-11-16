package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_untagResourceCmd).Standalone()

		wellarchitected_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		wellarchitected_untagResourceCmd.Flags().String("workload-arn", "", "")
		wellarchitected_untagResourceCmd.MarkFlagRequired("tag-keys")
		wellarchitected_untagResourceCmd.MarkFlagRequired("workload-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_untagResourceCmd)
}
