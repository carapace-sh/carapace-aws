package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_runInstancesCmd = &cobra.Command{
	Use:   "run-instances",
	Short: "Launches the specified number of instances using an AMI for which you have permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_runInstancesCmd).Standalone()

	ec2_runInstancesCmd.Flags().String("additional-info", "", "Reserved.")
	ec2_runInstancesCmd.Flags().String("block-device-mappings", "", "The block device mapping, which defines the EBS volumes and instance store volumes to attach to the instance at launch.")
	ec2_runInstancesCmd.Flags().String("capacity-reservation-specification", "", "Information about the Capacity Reservation targeting option.")
	ec2_runInstancesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
	ec2_runInstancesCmd.Flags().String("cpu-options", "", "The CPU options for the instance.")
	ec2_runInstancesCmd.Flags().String("credit-specification", "", "The credit option for CPU usage of the burstable performance instance.")
	ec2_runInstancesCmd.Flags().Bool("disable-api-stop", false, "Indicates whether an instance is enabled for stop protection.")
	ec2_runInstancesCmd.Flags().Bool("disable-api-termination", false, "Indicates whether termination protection is enabled for the instance.")
	ec2_runInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_runInstancesCmd.Flags().Bool("ebs-optimized", false, "Indicates whether the instance is optimized for Amazon EBS I/O.")
	ec2_runInstancesCmd.Flags().String("elastic-gpu-specification", "", "An elastic GPU to associate with the instance.")
	ec2_runInstancesCmd.Flags().String("elastic-inference-accelerators", "", "An elastic inference accelerator to associate with the instance.")
	ec2_runInstancesCmd.Flags().Bool("enable-primary-ipv6", false, "If you’re launching an instance into a dual-stack or IPv6-only subnet, you can enable assigning a primary IPv6 address.")
	ec2_runInstancesCmd.Flags().String("enclave-options", "", "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves.")
	ec2_runInstancesCmd.Flags().String("hibernation-options", "", "Indicates whether an instance is enabled for hibernation.")
	ec2_runInstancesCmd.Flags().String("iam-instance-profile", "", "The name or Amazon Resource Name (ARN) of an IAM instance profile.")
	ec2_runInstancesCmd.Flags().String("image-id", "", "The ID of the AMI.")
	ec2_runInstancesCmd.Flags().String("instance-initiated-shutdown-behavior", "", "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).")
	ec2_runInstancesCmd.Flags().String("instance-market-options", "", "The market (purchasing) option for the instances.")
	ec2_runInstancesCmd.Flags().String("instance-type", "", "The instance type.")
	ec2_runInstancesCmd.Flags().String("ipv6-address-count", "", "The number of IPv6 addresses to associate with the primary network interface.")
	ec2_runInstancesCmd.Flags().String("ipv6-addresses", "", "The IPv6 addresses from the range of the subnet to associate with the primary network interface.")
	ec2_runInstancesCmd.Flags().String("kernel-id", "", "The ID of the kernel.")
	ec2_runInstancesCmd.Flags().String("key-name", "", "The name of the key pair.")
	ec2_runInstancesCmd.Flags().String("launch-template", "", "The launch template.")
	ec2_runInstancesCmd.Flags().String("license-specifications", "", "The license configurations.")
	ec2_runInstancesCmd.Flags().String("maintenance-options", "", "The maintenance and recovery options for the instance.")
	ec2_runInstancesCmd.Flags().String("max-count", "", "The maximum number of instances to launch.")
	ec2_runInstancesCmd.Flags().String("metadata-options", "", "The metadata options for the instance.")
	ec2_runInstancesCmd.Flags().String("min-count", "", "The minimum number of instances to launch.")
	ec2_runInstancesCmd.Flags().String("monitoring", "", "Specifies whether detailed monitoring is enabled for the instance.")
	ec2_runInstancesCmd.Flags().String("network-interfaces", "", "The network interfaces to associate with the instance.")
	ec2_runInstancesCmd.Flags().String("network-performance-options", "", "Contains settings for the network performance options for the instance.")
	ec2_runInstancesCmd.Flags().Bool("no-disable-api-stop", false, "Indicates whether an instance is enabled for stop protection.")
	ec2_runInstancesCmd.Flags().Bool("no-disable-api-termination", false, "Indicates whether termination protection is enabled for the instance.")
	ec2_runInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_runInstancesCmd.Flags().Bool("no-ebs-optimized", false, "Indicates whether the instance is optimized for Amazon EBS I/O.")
	ec2_runInstancesCmd.Flags().Bool("no-enable-primary-ipv6", false, "If you’re launching an instance into a dual-stack or IPv6-only subnet, you can enable assigning a primary IPv6 address.")
	ec2_runInstancesCmd.Flags().String("operator", "", "Reserved for internal use.")
	ec2_runInstancesCmd.Flags().String("placement", "", "The placement for the instance.")
	ec2_runInstancesCmd.Flags().String("private-dns-name-options", "", "The options for the instance hostname.")
	ec2_runInstancesCmd.Flags().String("private-ip-address", "", "The primary IPv4 address.")
	ec2_runInstancesCmd.Flags().String("ramdisk-id", "", "The ID of the RAM disk to select.")
	ec2_runInstancesCmd.Flags().String("security-group-ids", "", "The IDs of the security groups.")
	ec2_runInstancesCmd.Flags().String("security-groups", "", "\\[Default VPC] The names of the security groups.")
	ec2_runInstancesCmd.Flags().String("subnet-id", "", "The ID of the subnet to launch the instance into.")
	ec2_runInstancesCmd.Flags().String("tag-specifications", "", "The tags to apply to the resources that are created during instance launch.")
	ec2_runInstancesCmd.Flags().String("user-data", "", "The user data to make available to the instance.")
	ec2_runInstancesCmd.MarkFlagRequired("max-count")
	ec2_runInstancesCmd.MarkFlagRequired("min-count")
	ec2_runInstancesCmd.Flag("no-disable-api-stop").Hidden = true
	ec2_runInstancesCmd.Flag("no-disable-api-termination").Hidden = true
	ec2_runInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2_runInstancesCmd.Flag("no-ebs-optimized").Hidden = true
	ec2_runInstancesCmd.Flag("no-enable-primary-ipv6").Hidden = true
	ec2Cmd.AddCommand(ec2_runInstancesCmd)
}
