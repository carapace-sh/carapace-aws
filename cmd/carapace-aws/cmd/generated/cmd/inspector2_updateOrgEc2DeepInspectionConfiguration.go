package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateOrgEc2DeepInspectionConfigurationCmd = &cobra.Command{
	Use:   "update-org-ec2-deep-inspection-configuration",
	Short: "Updates the Amazon Inspector deep inspection custom paths for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateOrgEc2DeepInspectionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_updateOrgEc2DeepInspectionConfigurationCmd).Standalone()

		inspector2_updateOrgEc2DeepInspectionConfigurationCmd.Flags().String("org-package-paths", "", "The Amazon Inspector deep inspection custom paths you are adding for your organization.")
		inspector2_updateOrgEc2DeepInspectionConfigurationCmd.MarkFlagRequired("org-package-paths")
	})
	inspector2Cmd.AddCommand(inspector2_updateOrgEc2DeepInspectionConfigurationCmd)
}
