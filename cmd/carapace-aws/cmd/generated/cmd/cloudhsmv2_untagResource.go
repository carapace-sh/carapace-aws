package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tag or tags from the specified CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_untagResourceCmd).Standalone()

	cloudhsmv2_untagResourceCmd.Flags().String("resource-id", "", "The cluster identifier (ID) for the cluster whose tags you are removing.")
	cloudhsmv2_untagResourceCmd.Flags().String("tag-key-list", "", "A list of one or more tag keys for the tags that you are removing.")
	cloudhsmv2_untagResourceCmd.MarkFlagRequired("resource-id")
	cloudhsmv2_untagResourceCmd.MarkFlagRequired("tag-key-list")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_untagResourceCmd)
}
