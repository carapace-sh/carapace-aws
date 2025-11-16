package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateConfigurationCmd = &cobra.Command{
	Use:   "update-configuration",
	Short: "Updates setting configurations for your Amazon Inspector account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateConfigurationCmd).Standalone()

	inspector2_updateConfigurationCmd.Flags().String("ec2-configuration", "", "Specifies how the Amazon EC2 automated scan will be updated for your environment.")
	inspector2_updateConfigurationCmd.Flags().String("ecr-configuration", "", "Specifies how the ECR automated re-scan will be updated for your environment.")
	inspector2Cmd.AddCommand(inspector2_updateConfigurationCmd)
}
