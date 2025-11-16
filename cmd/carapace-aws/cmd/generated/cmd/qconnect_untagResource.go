package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_untagResourceCmd).Standalone()

	qconnect_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	qconnect_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	qconnect_untagResourceCmd.MarkFlagRequired("resource-arn")
	qconnect_untagResourceCmd.MarkFlagRequired("tag-keys")
	qconnectCmd.AddCommand(qconnect_untagResourceCmd)
}
