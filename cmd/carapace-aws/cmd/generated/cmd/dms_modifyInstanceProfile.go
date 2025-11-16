package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyInstanceProfileCmd = &cobra.Command{
	Use:   "modify-instance-profile",
	Short: "Modifies the specified instance profile using the provided parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyInstanceProfileCmd).Standalone()

		dms_modifyInstanceProfileCmd.Flags().String("availability-zone", "", "The Availability Zone where the instance profile runs.")
		dms_modifyInstanceProfileCmd.Flags().String("description", "", "A user-friendly description for the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("instance-profile-identifier", "", "The identifier of the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("instance-profile-name", "", "A user-friendly name for the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key that is used to encrypt the connection parameters for the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("network-type", "", "Specifies the network type for the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("publicly-accessible", "", "Specifies the accessibility options for the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("subnet-group-identifier", "", "A subnet group to associate with the instance profile.")
		dms_modifyInstanceProfileCmd.Flags().String("vpc-security-groups", "", "Specifies the VPC security groups to be used with the instance profile.")
		dms_modifyInstanceProfileCmd.MarkFlagRequired("instance-profile-identifier")
	})
	dmsCmd.AddCommand(dms_modifyInstanceProfileCmd)
}
