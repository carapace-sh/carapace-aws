package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_registerImageCmd = &cobra.Command{
	Use:   "register-image",
	Short: "Registers an AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_registerImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_registerImageCmd).Standalone()

		ec2_registerImageCmd.Flags().String("architecture", "", "The architecture of the AMI.")
		ec2_registerImageCmd.Flags().String("billing-products", "", "The billing product codes.")
		ec2_registerImageCmd.Flags().String("block-device-mappings", "", "The block device mapping entries.")
		ec2_registerImageCmd.Flags().String("boot-mode", "", "The boot mode of the AMI.")
		ec2_registerImageCmd.Flags().String("description", "", "A description for your AMI.")
		ec2_registerImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_registerImageCmd.Flags().Bool("ena-support", false, "Set to `true` to enable enhanced networking with ENA for the AMI and any instances that you launch from the AMI.")
		ec2_registerImageCmd.Flags().String("image-location", "", "The full path to your AMI manifest in Amazon S3 storage.")
		ec2_registerImageCmd.Flags().String("imds-support", "", "Set to `v2.0` to indicate that IMDSv2 is specified in the AMI.")
		ec2_registerImageCmd.Flags().String("kernel-id", "", "The ID of the kernel.")
		ec2_registerImageCmd.Flags().String("name", "", "A name for your AMI.")
		ec2_registerImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_registerImageCmd.Flags().Bool("no-ena-support", false, "Set to `true` to enable enhanced networking with ENA for the AMI and any instances that you launch from the AMI.")
		ec2_registerImageCmd.Flags().String("ramdisk-id", "", "The ID of the RAM disk.")
		ec2_registerImageCmd.Flags().String("root-device-name", "", "The device name of the root device volume (for example, `/dev/sda1`).")
		ec2_registerImageCmd.Flags().String("sriov-net-support", "", "Set to `simple` to enable enhanced networking with the Intel 82599 Virtual Function interface for the AMI and any instances that you launch from the AMI.")
		ec2_registerImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the AMI.")
		ec2_registerImageCmd.Flags().String("tpm-support", "", "Set to `v2.0` to enable Trusted Platform Module (TPM) support.")
		ec2_registerImageCmd.Flags().String("uefi-data", "", "Base64 representation of the non-volatile UEFI variable store.")
		ec2_registerImageCmd.Flags().String("virtualization-type", "", "The type of virtualization (`hvm` | `paravirtual`).")
		ec2_registerImageCmd.MarkFlagRequired("name")
		ec2_registerImageCmd.Flag("no-dry-run").Hidden = true
		ec2_registerImageCmd.Flag("no-ena-support").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_registerImageCmd)
}
