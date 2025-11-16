package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_registerToWorkMailCmd = &cobra.Command{
	Use:   "register-to-work-mail",
	Short: "Registers an existing and disabled user, group, or resource for WorkMail use by associating a mailbox and calendaring capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_registerToWorkMailCmd).Standalone()

	workmail_registerToWorkMailCmd.Flags().String("email", "", "The email for the user, group, or resource to be updated.")
	workmail_registerToWorkMailCmd.Flags().String("entity-id", "", "The identifier for the user, group, or resource to be updated.")
	workmail_registerToWorkMailCmd.Flags().String("organization-id", "", "The identifier for the organization under which the user, group, or resource exists.")
	workmail_registerToWorkMailCmd.MarkFlagRequired("email")
	workmail_registerToWorkMailCmd.MarkFlagRequired("entity-id")
	workmail_registerToWorkMailCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_registerToWorkMailCmd)
}
