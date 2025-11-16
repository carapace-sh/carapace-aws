package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listImpersonationRolesCmd = &cobra.Command{
	Use:   "list-impersonation-roles",
	Short: "Lists all the impersonation roles for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listImpersonationRolesCmd).Standalone()

	workmail_listImpersonationRolesCmd.Flags().String("max-results", "", "The maximum number of results returned in a single call.")
	workmail_listImpersonationRolesCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
	workmail_listImpersonationRolesCmd.Flags().String("organization-id", "", "The WorkMail organization to which the listed impersonation roles belong.")
	workmail_listImpersonationRolesCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listImpersonationRolesCmd)
}
