package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an IoT Device Advisor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_untagResourceCmd).Standalone()

	iotdeviceadvisor_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN of an IoT Device Advisor resource.")
	iotdeviceadvisor_untagResourceCmd.Flags().String("tag-keys", "", "List of tag keys to remove from the IoT Device Advisor resource.")
	iotdeviceadvisor_untagResourceCmd.MarkFlagRequired("resource-arn")
	iotdeviceadvisor_untagResourceCmd.MarkFlagRequired("tag-keys")
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_untagResourceCmd)
}
