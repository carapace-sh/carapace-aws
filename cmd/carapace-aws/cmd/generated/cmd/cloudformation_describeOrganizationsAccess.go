package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeOrganizationsAccessCmd = &cobra.Command{
	Use:   "describe-organizations-access",
	Short: "Retrieves information about the account's `OrganizationAccess` status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeOrganizationsAccessCmd).Standalone()

	cloudformation_describeOrganizationsAccessCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformationCmd.AddCommand(cloudformation_describeOrganizationsAccessCmd)
}
