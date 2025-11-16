package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createInstanceProfileCmd = &cobra.Command{
	Use:   "create-instance-profile",
	Short: "Creates the instance profile using the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createInstanceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createInstanceProfileCmd).Standalone()

		dms_createInstanceProfileCmd.Flags().String("availability-zone", "", "The Availability Zone where the instance profile will be created.")
		dms_createInstanceProfileCmd.Flags().String("description", "", "A user-friendly description of the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("instance-profile-name", "", "A user-friendly name for the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key that is used to encrypt the connection parameters for the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("network-type", "", "Specifies the network type for the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("publicly-accessible", "", "Specifies the accessibility options for the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("subnet-group-identifier", "", "A subnet group to associate with the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("tags", "", "One or more tags to be assigned to the instance profile.")
		dms_createInstanceProfileCmd.Flags().String("vpc-security-groups", "", "Specifies the VPC security group names to be used with the instance profile.")
	})
	dmsCmd.AddCommand(dms_createInstanceProfileCmd)
}
