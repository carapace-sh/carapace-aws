package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getEc2DeepInspectionConfigurationCmd = &cobra.Command{
	Use:   "get-ec2-deep-inspection-configuration",
	Short: "Retrieves the activation status of Amazon Inspector deep inspection and custom paths associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getEc2DeepInspectionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getEc2DeepInspectionConfigurationCmd).Standalone()

	})
	inspector2Cmd.AddCommand(inspector2_getEc2DeepInspectionConfigurationCmd)
}
