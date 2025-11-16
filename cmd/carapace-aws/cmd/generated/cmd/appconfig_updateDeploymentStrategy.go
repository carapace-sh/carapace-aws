package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateDeploymentStrategyCmd = &cobra.Command{
	Use:   "update-deployment-strategy",
	Short: "Updates a deployment strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateDeploymentStrategyCmd).Standalone()

	appconfig_updateDeploymentStrategyCmd.Flags().String("deployment-duration-in-minutes", "", "Total amount of time for a deployment to last.")
	appconfig_updateDeploymentStrategyCmd.Flags().String("deployment-strategy-id", "", "The deployment strategy ID.")
	appconfig_updateDeploymentStrategyCmd.Flags().String("description", "", "A description of the deployment strategy.")
	appconfig_updateDeploymentStrategyCmd.Flags().String("final-bake-time-in-minutes", "", "The amount of time that AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic rollback.")
	appconfig_updateDeploymentStrategyCmd.Flags().String("growth-factor", "", "The percentage of targets to receive a deployed configuration during each interval.")
	appconfig_updateDeploymentStrategyCmd.Flags().String("growth-type", "", "The algorithm used to define how percentage grows over time.")
	appconfig_updateDeploymentStrategyCmd.MarkFlagRequired("deployment-strategy-id")
	appconfigCmd.AddCommand(appconfig_updateDeploymentStrategyCmd)
}
