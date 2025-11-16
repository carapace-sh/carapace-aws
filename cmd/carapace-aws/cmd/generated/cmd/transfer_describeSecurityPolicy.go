package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeSecurityPolicyCmd = &cobra.Command{
	Use:   "describe-security-policy",
	Short: "Describes the security policy that is attached to your server or SFTP connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeSecurityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeSecurityPolicyCmd).Standalone()

		transfer_describeSecurityPolicyCmd.Flags().String("security-policy-name", "", "Specify the text name of the security policy for which you want the details.")
		transfer_describeSecurityPolicyCmd.MarkFlagRequired("security-policy-name")
	})
	transferCmd.AddCommand(transfer_describeSecurityPolicyCmd)
}
