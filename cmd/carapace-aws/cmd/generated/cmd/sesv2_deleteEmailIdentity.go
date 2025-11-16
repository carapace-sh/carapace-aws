package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteEmailIdentityCmd = &cobra.Command{
	Use:   "delete-email-identity",
	Short: "Deletes an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteEmailIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_deleteEmailIdentityCmd).Standalone()

		sesv2_deleteEmailIdentityCmd.Flags().String("email-identity", "", "The identity (that is, the email address or domain) to delete.")
		sesv2_deleteEmailIdentityCmd.MarkFlagRequired("email-identity")
	})
	sesv2Cmd.AddCommand(sesv2_deleteEmailIdentityCmd)
}
