package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunneling_untagResourceCmd).Standalone()

		iotsecuretunneling_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		iotsecuretunneling_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove.")
		iotsecuretunneling_untagResourceCmd.MarkFlagRequired("resource-arn")
		iotsecuretunneling_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_untagResourceCmd)
}
