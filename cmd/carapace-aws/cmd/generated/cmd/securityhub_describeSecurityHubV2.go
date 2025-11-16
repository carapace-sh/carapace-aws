package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeSecurityHubV2Cmd = &cobra.Command{
	Use:   "describe-security-hub-v2",
	Short: "Returns details about the service resource in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeSecurityHubV2Cmd).Standalone()

	securityhubCmd.AddCommand(securityhub_describeSecurityHubV2Cmd)
}
