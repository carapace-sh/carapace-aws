package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deassigns a resource tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_untagResourceCmd).Standalone()

	groundstation_untagResourceCmd.Flags().String("resource-arn", "", "ARN of a resource.")
	groundstation_untagResourceCmd.Flags().String("tag-keys", "", "Keys of a resource tag.")
	groundstation_untagResourceCmd.MarkFlagRequired("resource-arn")
	groundstation_untagResourceCmd.MarkFlagRequired("tag-keys")
	groundstationCmd.AddCommand(groundstation_untagResourceCmd)
}
