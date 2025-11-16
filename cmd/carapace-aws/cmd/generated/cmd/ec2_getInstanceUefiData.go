package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getInstanceUefiDataCmd = &cobra.Command{
	Use:   "get-instance-uefi-data",
	Short: "A binary representation of the UEFI variable store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getInstanceUefiDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getInstanceUefiDataCmd).Standalone()

		ec2_getInstanceUefiDataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getInstanceUefiDataCmd.Flags().String("instance-id", "", "The ID of the instance from which to retrieve the UEFI data.")
		ec2_getInstanceUefiDataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getInstanceUefiDataCmd.MarkFlagRequired("instance-id")
		ec2_getInstanceUefiDataCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getInstanceUefiDataCmd)
}
