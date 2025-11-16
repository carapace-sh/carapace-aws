package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeSecurityConfigurationCmd = &cobra.Command{
	Use:   "describe-security-configuration",
	Short: "Provides the details of a security configuration by returning the configuration JSON.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_describeSecurityConfigurationCmd).Standalone()

		emr_describeSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration.")
		emr_describeSecurityConfigurationCmd.MarkFlagRequired("name")
	})
	emrCmd.AddCommand(emr_describeSecurityConfigurationCmd)
}
