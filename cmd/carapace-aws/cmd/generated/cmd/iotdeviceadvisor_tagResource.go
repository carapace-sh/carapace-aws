package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to and modifies existing tags of an IoT Device Advisor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_tagResourceCmd).Standalone()

	iotdeviceadvisor_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN of an IoT Device Advisor resource.")
	iotdeviceadvisor_tagResourceCmd.Flags().String("tags", "", "The tags to be attached to the IoT Device Advisor resource.")
	iotdeviceadvisor_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotdeviceadvisor_tagResourceCmd.MarkFlagRequired("tags")
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_tagResourceCmd)
}
