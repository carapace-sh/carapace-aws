package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setIdentityDkimEnabledCmd = &cobra.Command{
	Use:   "set-identity-dkim-enabled",
	Short: "Enables or disables Easy DKIM signing of email sent from an identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setIdentityDkimEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_setIdentityDkimEnabledCmd).Standalone()

		ses_setIdentityDkimEnabledCmd.Flags().String("dkim-enabled", "", "Sets whether DKIM signing is enabled for an identity.")
		ses_setIdentityDkimEnabledCmd.Flags().String("identity", "", "The identity for which DKIM signing should be enabled or disabled.")
		ses_setIdentityDkimEnabledCmd.MarkFlagRequired("dkim-enabled")
		ses_setIdentityDkimEnabledCmd.MarkFlagRequired("identity")
	})
	sesCmd.AddCommand(ses_setIdentityDkimEnabledCmd)
}
