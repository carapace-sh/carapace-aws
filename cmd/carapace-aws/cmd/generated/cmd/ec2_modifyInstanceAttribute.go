package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceAttributeCmd = &cobra.Command{
	Use:   "modify-instance-attribute",
	Short: "Modifies the specified attribute of the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceAttributeCmd).Standalone()

	ec2_modifyInstanceAttributeCmd.Flags().String("attribute", "", "The name of the attribute to modify.")
	ec2_modifyInstanceAttributeCmd.Flags().String("block-device-mappings", "", "Modifies the `DeleteOnTermination` attribute for volumes that are currently attached.")
	ec2_modifyInstanceAttributeCmd.Flags().String("disable-api-stop", "", "Indicates whether an instance is enabled for stop protection.")
	ec2_modifyInstanceAttributeCmd.Flags().String("disable-api-termination", "", "Enable or disable termination protection for the instance.")
	ec2_modifyInstanceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceAttributeCmd.Flags().String("ebs-optimized", "", "Specifies whether the instance is optimized for Amazon EBS I/O.")
	ec2_modifyInstanceAttributeCmd.Flags().String("ena-support", "", "Set to `true` to enable enhanced networking with ENA for the instance.")
	ec2_modifyInstanceAttributeCmd.Flags().String("groups", "", "Replaces the security groups of the instance with the specified security groups.")
	ec2_modifyInstanceAttributeCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_modifyInstanceAttributeCmd.Flags().String("instance-initiated-shutdown-behavior", "", "Specifies whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).")
	ec2_modifyInstanceAttributeCmd.Flags().String("instance-type", "", "Changes the instance type to the specified value.")
	ec2_modifyInstanceAttributeCmd.Flags().String("kernel", "", "Changes the instance's kernel to the specified value.")
	ec2_modifyInstanceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceAttributeCmd.Flags().String("ramdisk", "", "Changes the instance's RAM disk to the specified value.")
	ec2_modifyInstanceAttributeCmd.Flags().String("source-dest-check", "", "Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.")
	ec2_modifyInstanceAttributeCmd.Flags().String("sriov-net-support", "", "Set to `simple` to enable enhanced networking with the Intel 82599 Virtual Function interface for the instance.")
	ec2_modifyInstanceAttributeCmd.Flags().String("user-data", "", "Changes the instance's user data to the specified value.")
	ec2_modifyInstanceAttributeCmd.Flags().String("value", "", "A new value for the attribute.")
	ec2_modifyInstanceAttributeCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyInstanceAttributeCmd)
}
