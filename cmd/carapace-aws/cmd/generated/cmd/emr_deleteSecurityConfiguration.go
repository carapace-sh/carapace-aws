package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_deleteSecurityConfigurationCmd = &cobra.Command{
	Use:   "delete-security-configuration",
	Short: "Deletes a security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_deleteSecurityConfigurationCmd).Standalone()

	emr_deleteSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration.")
	emr_deleteSecurityConfigurationCmd.MarkFlagRequired("name")
	emrCmd.AddCommand(emr_deleteSecurityConfigurationCmd)
}
