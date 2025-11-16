package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates tags from an Amazon Q Apps resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_untagResourceCmd).Standalone()

	qapps_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to disassociate the tag from.")
	qapps_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to disassociate from the resource.")
	qapps_untagResourceCmd.MarkFlagRequired("resource-arn")
	qapps_untagResourceCmd.MarkFlagRequired("tag-keys")
	qappsCmd.AddCommand(qapps_untagResourceCmd)
}
