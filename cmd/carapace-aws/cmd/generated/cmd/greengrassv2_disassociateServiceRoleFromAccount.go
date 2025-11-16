package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_disassociateServiceRoleFromAccountCmd = &cobra.Command{
	Use:   "disassociate-service-role-from-account",
	Short: "Disassociates the Greengrass service role from IoT Greengrass for your Amazon Web Services account in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_disassociateServiceRoleFromAccountCmd).Standalone()

	greengrassv2Cmd.AddCommand(greengrassv2_disassociateServiceRoleFromAccountCmd)
}
