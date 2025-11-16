package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from an Amazon EVS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_untagResourceCmd).Standalone()

	evs_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to delete tags from.")
	evs_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to delete.")
	evs_untagResourceCmd.MarkFlagRequired("resource-arn")
	evs_untagResourceCmd.MarkFlagRequired("tag-keys")
	evsCmd.AddCommand(evs_untagResourceCmd)
}
