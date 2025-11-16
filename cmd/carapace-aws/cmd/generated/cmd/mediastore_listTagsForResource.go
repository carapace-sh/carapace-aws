package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags assigned to the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_listTagsForResourceCmd).Standalone()

		mediastore_listTagsForResourceCmd.Flags().String("resource", "", "The Amazon Resource Name (ARN) for the container.")
		mediastore_listTagsForResourceCmd.MarkFlagRequired("resource")
	})
	mediastoreCmd.AddCommand(mediastore_listTagsForResourceCmd)
}
