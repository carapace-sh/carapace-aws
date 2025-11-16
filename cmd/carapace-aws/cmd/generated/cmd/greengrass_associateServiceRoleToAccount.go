package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_associateServiceRoleToAccountCmd = &cobra.Command{
	Use:   "associate-service-role-to-account",
	Short: "Associates a role with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_associateServiceRoleToAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_associateServiceRoleToAccountCmd).Standalone()

		greengrass_associateServiceRoleToAccountCmd.Flags().String("role-arn", "", "The ARN of the service role you wish to associate with your account.")
		greengrass_associateServiceRoleToAccountCmd.MarkFlagRequired("role-arn")
	})
	greengrassCmd.AddCommand(greengrass_associateServiceRoleToAccountCmd)
}
