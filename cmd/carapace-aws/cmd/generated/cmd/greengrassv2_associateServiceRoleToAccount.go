package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_associateServiceRoleToAccountCmd = &cobra.Command{
	Use:   "associate-service-role-to-account",
	Short: "Associates a Greengrass service role with IoT Greengrass for your Amazon Web Services account in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_associateServiceRoleToAccountCmd).Standalone()

	greengrassv2_associateServiceRoleToAccountCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the service role to associate with IoT Greengrass for your Amazon Web Services account in this Amazon Web Services Region.")
	greengrassv2_associateServiceRoleToAccountCmd.MarkFlagRequired("role-arn")
	greengrassv2Cmd.AddCommand(greengrassv2_associateServiceRoleToAccountCmd)
}
