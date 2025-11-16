package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_tagResourceCmd).Standalone()

		wellarchitected_tagResourceCmd.Flags().String("tags", "", "The tags for the resource.")
		wellarchitected_tagResourceCmd.Flags().String("workload-arn", "", "")
		wellarchitected_tagResourceCmd.MarkFlagRequired("tags")
		wellarchitected_tagResourceCmd.MarkFlagRequired("workload-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_tagResourceCmd)
}
