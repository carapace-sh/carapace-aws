package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "delete-launch-configuration-template",
	Short: "Deletes a single Launch Configuration Template by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteLaunchConfigurationTemplateCmd).Standalone()

	mgn_deleteLaunchConfigurationTemplateCmd.Flags().String("launch-configuration-template-id", "", "ID of resource to be deleted.")
	mgn_deleteLaunchConfigurationTemplateCmd.MarkFlagRequired("launch-configuration-template-id")
	mgnCmd.AddCommand(mgn_deleteLaunchConfigurationTemplateCmd)
}
