package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdminCmd = &cobra.Command{
	Use:   "sso-admin",
	Short: "IAM Identity Center is the Amazon Web Services solution for connecting your workforce users to Amazon Web Services managed applications and other Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdminCmd).Standalone()

	rootCmd.AddCommand(ssoAdminCmd)
}
