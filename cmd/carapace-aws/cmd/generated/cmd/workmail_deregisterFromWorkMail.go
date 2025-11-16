package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deregisterFromWorkMailCmd = &cobra.Command{
	Use:   "deregister-from-work-mail",
	Short: "Mark a user, group, or resource as no longer used in WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deregisterFromWorkMailCmd).Standalone()

	workmail_deregisterFromWorkMailCmd.Flags().String("entity-id", "", "The identifier for the member to be updated.")
	workmail_deregisterFromWorkMailCmd.Flags().String("organization-id", "", "The identifier for the organization under which the WorkMail entity exists.")
	workmail_deregisterFromWorkMailCmd.MarkFlagRequired("entity-id")
	workmail_deregisterFromWorkMailCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deregisterFromWorkMailCmd)
}
