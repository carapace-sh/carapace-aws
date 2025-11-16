package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getEmailIdentityCmd = &cobra.Command{
	Use:   "get-email-identity",
	Short: "Provides information about a specific identity, including the identity's verification status, sending authorization policies, its DKIM authentication status, and its custom Mail-From settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getEmailIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getEmailIdentityCmd).Standalone()

		sesv2_getEmailIdentityCmd.Flags().String("email-identity", "", "The email identity.")
		sesv2_getEmailIdentityCmd.MarkFlagRequired("email-identity")
	})
	sesv2Cmd.AddCommand(sesv2_getEmailIdentityCmd)
}
