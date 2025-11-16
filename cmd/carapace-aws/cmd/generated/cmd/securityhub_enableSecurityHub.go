package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_enableSecurityHubCmd = &cobra.Command{
	Use:   "enable-security-hub",
	Short: "Enables Security Hub for your account in the current Region or the Region you specify in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_enableSecurityHubCmd).Standalone()

	securityhub_enableSecurityHubCmd.Flags().String("control-finding-generator", "", "This field, used when enabling Security Hub, specifies whether the calling account has consolidated control findings turned on.")
	securityhub_enableSecurityHubCmd.Flags().Bool("enable-default-standards", false, "Whether to enable the security standards that Security Hub has designated as automatically enabled.")
	securityhub_enableSecurityHubCmd.Flags().Bool("no-enable-default-standards", false, "Whether to enable the security standards that Security Hub has designated as automatically enabled.")
	securityhub_enableSecurityHubCmd.Flags().String("tags", "", "The tags to add to the hub resource when you enable Security Hub.")
	securityhub_enableSecurityHubCmd.Flag("no-enable-default-standards").Hidden = true
	securityhubCmd.AddCommand(securityhub_enableSecurityHubCmd)
}
