package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_describeDeviceEc2InstancesCmd = &cobra.Command{
	Use:   "describe-device-ec2-instances",
	Short: "Checks the current state of the Amazon EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_describeDeviceEc2InstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_describeDeviceEc2InstancesCmd).Standalone()

		snowDeviceManagement_describeDeviceEc2InstancesCmd.Flags().String("instance-ids", "", "A list of instance IDs associated with the managed device.")
		snowDeviceManagement_describeDeviceEc2InstancesCmd.Flags().String("managed-device-id", "", "The ID of the managed device.")
		snowDeviceManagement_describeDeviceEc2InstancesCmd.MarkFlagRequired("instance-ids")
		snowDeviceManagement_describeDeviceEc2InstancesCmd.MarkFlagRequired("managed-device-id")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_describeDeviceEc2InstancesCmd)
}
