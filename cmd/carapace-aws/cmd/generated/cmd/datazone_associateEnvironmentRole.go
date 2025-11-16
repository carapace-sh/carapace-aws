package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_associateEnvironmentRoleCmd = &cobra.Command{
	Use:   "associate-environment-role",
	Short: "Associates the environment role in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_associateEnvironmentRoleCmd).Standalone()

	datazone_associateEnvironmentRoleCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the environment role is associated.")
	datazone_associateEnvironmentRoleCmd.Flags().String("environment-identifier", "", "The ID of the Amazon DataZone environment.")
	datazone_associateEnvironmentRoleCmd.Flags().String("environment-role-arn", "", "The ARN of the environment role.")
	datazone_associateEnvironmentRoleCmd.MarkFlagRequired("domain-identifier")
	datazone_associateEnvironmentRoleCmd.MarkFlagRequired("environment-identifier")
	datazone_associateEnvironmentRoleCmd.MarkFlagRequired("environment-role-arn")
	datazoneCmd.AddCommand(datazone_associateEnvironmentRoleCmd)
}
