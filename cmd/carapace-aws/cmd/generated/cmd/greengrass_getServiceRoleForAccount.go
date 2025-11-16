package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getServiceRoleForAccountCmd = &cobra.Command{
	Use:   "get-service-role-for-account",
	Short: "Retrieves the service role that is attached to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getServiceRoleForAccountCmd).Standalone()

	greengrassCmd.AddCommand(greengrass_getServiceRoleForAccountCmd)
}
