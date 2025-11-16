package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_createLaunchConfigurationCmd = &cobra.Command{
	Use:   "create-launch-configuration",
	Short: "Creates a launch configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_createLaunchConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_createLaunchConfigurationCmd).Standalone()

		autoscaling_createLaunchConfigurationCmd.Flags().String("associate-public-ip-address", "", "Specifies whether to assign a public IPv4 address to the group's instances.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("block-device-mappings", "", "The block device mapping entries that define the block devices to attach to the instances at launch.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("classic-link-vpcid", "", "Available for backward compatibility.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("classic-link-vpcsecurity-groups", "", "Available for backward compatibility.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("ebs-optimized", "", "Specifies whether the launch configuration is optimized for EBS I/O (`true`) or not (`false`).")
		autoscaling_createLaunchConfigurationCmd.Flags().String("iam-instance-profile", "", "The name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("image-id", "", "The ID of the Amazon Machine Image (AMI) that was assigned during registration.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("instance-id", "", "The ID of the instance to use to create the launch configuration.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("instance-monitoring", "", "Controls whether instances in this group are launched with detailed (`true`) or basic (`false`) monitoring.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("instance-type", "", "Specifies the instance type of the EC2 instance.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("kernel-id", "", "The ID of the kernel associated with the AMI.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("key-name", "", "The name of the key pair.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("launch-configuration-name", "", "The name of the launch configuration.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("metadata-options", "", "The metadata options for the instances.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("placement-tenancy", "", "The tenancy of the instance, either `default` or `dedicated`.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("ramdisk-id", "", "The ID of the RAM disk to select.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("security-groups", "", "A list that contains the security group IDs to assign to the instances in the Auto Scaling group.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("spot-price", "", "The maximum hourly price to be paid for any Spot Instance launched to fulfill the request.")
		autoscaling_createLaunchConfigurationCmd.Flags().String("user-data", "", "The user data to make available to the launched EC2 instances.")
		autoscaling_createLaunchConfigurationCmd.MarkFlagRequired("launch-configuration-name")
	})
	autoscalingCmd.AddCommand(autoscaling_createLaunchConfigurationCmd)
}
