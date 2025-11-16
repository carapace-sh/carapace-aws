package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates the configuration details for a specific runtime environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_updateEnvironmentCmd).Standalone()

	m2_updateEnvironmentCmd.Flags().Bool("apply-during-maintenance-window", false, "Indicates whether to update the runtime environment during the maintenance window.")
	m2_updateEnvironmentCmd.Flags().String("desired-capacity", "", "The desired capacity for the runtime environment to update.")
	m2_updateEnvironmentCmd.Flags().String("engine-version", "", "The version of the runtime engine for the runtime environment.")
	m2_updateEnvironmentCmd.Flags().String("environment-id", "", "The unique identifier of the runtime environment that you want to update.")
	m2_updateEnvironmentCmd.Flags().Bool("force-update", false, "Forces the updates on the environment.")
	m2_updateEnvironmentCmd.Flags().String("instance-type", "", "The instance type for the runtime environment to update.")
	m2_updateEnvironmentCmd.Flags().Bool("no-apply-during-maintenance-window", false, "Indicates whether to update the runtime environment during the maintenance window.")
	m2_updateEnvironmentCmd.Flags().Bool("no-force-update", false, "Forces the updates on the environment.")
	m2_updateEnvironmentCmd.Flags().String("preferred-maintenance-window", "", "Configures the maintenance window that you want for the runtime environment.")
	m2_updateEnvironmentCmd.MarkFlagRequired("environment-id")
	m2_updateEnvironmentCmd.Flag("no-apply-during-maintenance-window").Hidden = true
	m2_updateEnvironmentCmd.Flag("no-force-update").Hidden = true
	m2Cmd.AddCommand(m2_updateEnvironmentCmd)
}
