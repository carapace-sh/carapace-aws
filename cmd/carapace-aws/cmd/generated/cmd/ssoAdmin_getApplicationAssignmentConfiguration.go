package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getApplicationAssignmentConfigurationCmd = &cobra.Command{
	Use:   "get-application-assignment-configuration",
	Short: "Retrieves the configuration of [PutApplicationAssignmentConfiguration]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getApplicationAssignmentConfigurationCmd).Standalone()

	ssoAdmin_getApplicationAssignmentConfigurationCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
	ssoAdmin_getApplicationAssignmentConfigurationCmd.MarkFlagRequired("application-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_getApplicationAssignmentConfigurationCmd)
}
