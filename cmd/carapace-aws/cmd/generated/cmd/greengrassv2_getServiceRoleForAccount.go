package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getServiceRoleForAccountCmd = &cobra.Command{
	Use:   "get-service-role-for-account",
	Short: "Gets the service role associated with IoT Greengrass for your Amazon Web Services account in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getServiceRoleForAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_getServiceRoleForAccountCmd).Standalone()

	})
	greengrassv2Cmd.AddCommand(greengrassv2_getServiceRoleForAccountCmd)
}
