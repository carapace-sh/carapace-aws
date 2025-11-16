package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateEc2DeepInspectionConfigurationCmd = &cobra.Command{
	Use:   "update-ec2-deep-inspection-configuration",
	Short: "Activates, deactivates Amazon Inspector deep inspection, or updates custom paths for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateEc2DeepInspectionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_updateEc2DeepInspectionConfigurationCmd).Standalone()

		inspector2_updateEc2DeepInspectionConfigurationCmd.Flags().Bool("activate-deep-inspection", false, "Specify `TRUE` to activate Amazon Inspector deep inspection in your account, or `FALSE` to deactivate.")
		inspector2_updateEc2DeepInspectionConfigurationCmd.Flags().Bool("no-activate-deep-inspection", false, "Specify `TRUE` to activate Amazon Inspector deep inspection in your account, or `FALSE` to deactivate.")
		inspector2_updateEc2DeepInspectionConfigurationCmd.Flags().String("package-paths", "", "The Amazon Inspector deep inspection custom paths you are adding for your account.")
		inspector2_updateEc2DeepInspectionConfigurationCmd.Flag("no-activate-deep-inspection").Hidden = true
	})
	inspector2Cmd.AddCommand(inspector2_updateEc2DeepInspectionConfigurationCmd)
}
