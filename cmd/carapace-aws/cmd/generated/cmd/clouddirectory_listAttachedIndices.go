package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listAttachedIndicesCmd = &cobra.Command{
	Use:   "list-attached-indices",
	Short: "Lists indices attached to the specified object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listAttachedIndicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listAttachedIndicesCmd).Standalone()

		clouddirectory_listAttachedIndicesCmd.Flags().String("consistency-level", "", "The consistency level to use for this operation.")
		clouddirectory_listAttachedIndicesCmd.Flags().String("directory-arn", "", "The ARN of the directory.")
		clouddirectory_listAttachedIndicesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listAttachedIndicesCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listAttachedIndicesCmd.Flags().String("target-reference", "", "A reference to the object that has indices attached.")
		clouddirectory_listAttachedIndicesCmd.MarkFlagRequired("directory-arn")
		clouddirectory_listAttachedIndicesCmd.MarkFlagRequired("target-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listAttachedIndicesCmd)
}
