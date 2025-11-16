package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "delete-launch-configuration-template",
	Short: "Deletes a single Launch Configuration Template by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteLaunchConfigurationTemplateCmd).Standalone()

	drs_deleteLaunchConfigurationTemplateCmd.Flags().String("launch-configuration-template-id", "", "The ID of the Launch Configuration Template to be deleted.")
	drs_deleteLaunchConfigurationTemplateCmd.MarkFlagRequired("launch-configuration-template-id")
	drsCmd.AddCommand(drs_deleteLaunchConfigurationTemplateCmd)
}
