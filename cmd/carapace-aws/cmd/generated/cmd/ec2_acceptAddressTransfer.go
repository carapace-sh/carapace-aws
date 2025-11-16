package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptAddressTransferCmd = &cobra.Command{
	Use:   "accept-address-transfer",
	Short: "Accepts an Elastic IP address transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptAddressTransferCmd).Standalone()

	ec2_acceptAddressTransferCmd.Flags().String("address", "", "The Elastic IP address you are accepting for transfer.")
	ec2_acceptAddressTransferCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptAddressTransferCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptAddressTransferCmd.Flags().String("tag-specifications", "", "`tag`:&lt;key&gt; - The key/value combination of a tag assigned to the resource.")
	ec2_acceptAddressTransferCmd.MarkFlagRequired("address")
	ec2_acceptAddressTransferCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_acceptAddressTransferCmd)
}
