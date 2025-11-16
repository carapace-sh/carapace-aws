package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the given tags (metadata) from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_untagResourceCmd).Standalone()

	iot_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iot_untagResourceCmd.Flags().String("tag-keys", "", "A list of the keys of the tags to be removed from the resource.")
	iot_untagResourceCmd.MarkFlagRequired("resource-arn")
	iot_untagResourceCmd.MarkFlagRequired("tag-keys")
	iotCmd.AddCommand(iot_untagResourceCmd)
}
