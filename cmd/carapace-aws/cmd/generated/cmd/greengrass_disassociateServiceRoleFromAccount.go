package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_disassociateServiceRoleFromAccountCmd = &cobra.Command{
	Use:   "disassociate-service-role-from-account",
	Short: "Disassociates the service role from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_disassociateServiceRoleFromAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_disassociateServiceRoleFromAccountCmd).Standalone()

	})
	greengrassCmd.AddCommand(greengrass_disassociateServiceRoleFromAccountCmd)
}
