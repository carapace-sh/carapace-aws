package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_disassociateEnvironmentRoleCmd = &cobra.Command{
	Use:   "disassociate-environment-role",
	Short: "Disassociates the environment role in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_disassociateEnvironmentRoleCmd).Standalone()

	datazone_disassociateEnvironmentRoleCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which an environment role is disassociated.")
	datazone_disassociateEnvironmentRoleCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
	datazone_disassociateEnvironmentRoleCmd.Flags().String("environment-role-arn", "", "The ARN of the environment role.")
	datazone_disassociateEnvironmentRoleCmd.MarkFlagRequired("domain-identifier")
	datazone_disassociateEnvironmentRoleCmd.MarkFlagRequired("environment-identifier")
	datazone_disassociateEnvironmentRoleCmd.MarkFlagRequired("environment-role-arn")
	datazoneCmd.AddCommand(datazone_disassociateEnvironmentRoleCmd)
}
