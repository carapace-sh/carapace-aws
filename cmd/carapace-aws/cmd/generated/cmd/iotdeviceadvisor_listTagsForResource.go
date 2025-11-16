package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags attached to an IoT Device Advisor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_listTagsForResourceCmd).Standalone()

		iotdeviceadvisor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN of the IoT Device Advisor resource.")
		iotdeviceadvisor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_listTagsForResourceCmd)
}
