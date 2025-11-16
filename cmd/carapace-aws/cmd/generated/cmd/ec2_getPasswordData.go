package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getPasswordDataCmd = &cobra.Command{
	Use:   "get-password-data",
	Short: "Retrieves the encrypted administrator password for a running Windows instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getPasswordDataCmd).Standalone()

	ec2_getPasswordDataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_getPasswordDataCmd.Flags().String("instance-id", "", "The ID of the Windows instance.")
	ec2_getPasswordDataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_getPasswordDataCmd.MarkFlagRequired("instance-id")
	ec2_getPasswordDataCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getPasswordDataCmd)
}
