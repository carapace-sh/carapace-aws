package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_restoreAddressToClassicCmd = &cobra.Command{
	Use:   "restore-address-to-classic",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_restoreAddressToClassicCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_restoreAddressToClassicCmd).Standalone()

		ec2_restoreAddressToClassicCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreAddressToClassicCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreAddressToClassicCmd.Flags().String("public-ip", "", "The Elastic IP address.")
		ec2_restoreAddressToClassicCmd.Flag("no-dry-run").Hidden = true
		ec2_restoreAddressToClassicCmd.MarkFlagRequired("public-ip")
	})
	ec2Cmd.AddCommand(ec2_restoreAddressToClassicCmd)
}
