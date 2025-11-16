package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the specified tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_deleteTagsCmd).Standalone()

		autoscaling_deleteTagsCmd.Flags().String("tags", "", "One or more tags.")
		autoscaling_deleteTagsCmd.MarkFlagRequired("tags")
	})
	autoscalingCmd.AddCommand(autoscaling_deleteTagsCmd)
}
