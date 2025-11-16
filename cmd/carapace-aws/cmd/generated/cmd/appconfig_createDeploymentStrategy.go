package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createDeploymentStrategyCmd = &cobra.Command{
	Use:   "create-deployment-strategy",
	Short: "Creates a deployment strategy that defines important criteria for rolling out your configuration to the designated targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createDeploymentStrategyCmd).Standalone()

	appconfig_createDeploymentStrategyCmd.Flags().String("deployment-duration-in-minutes", "", "Total amount of time for a deployment to last.")
	appconfig_createDeploymentStrategyCmd.Flags().String("description", "", "A description of the deployment strategy.")
	appconfig_createDeploymentStrategyCmd.Flags().String("final-bake-time-in-minutes", "", "Specifies the amount of time AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete.")
	appconfig_createDeploymentStrategyCmd.Flags().String("growth-factor", "", "The percentage of targets to receive a deployed configuration during each interval.")
	appconfig_createDeploymentStrategyCmd.Flags().String("growth-type", "", "The algorithm used to define how percentage grows over time.")
	appconfig_createDeploymentStrategyCmd.Flags().String("name", "", "A name for the deployment strategy.")
	appconfig_createDeploymentStrategyCmd.Flags().String("replicate-to", "", "Save the deployment strategy to a Systems Manager (SSM) document.")
	appconfig_createDeploymentStrategyCmd.Flags().String("tags", "", "Metadata to assign to the deployment strategy.")
	appconfig_createDeploymentStrategyCmd.MarkFlagRequired("deployment-duration-in-minutes")
	appconfig_createDeploymentStrategyCmd.MarkFlagRequired("growth-factor")
	appconfig_createDeploymentStrategyCmd.MarkFlagRequired("name")
	appconfigCmd.AddCommand(appconfig_createDeploymentStrategyCmd)
}
