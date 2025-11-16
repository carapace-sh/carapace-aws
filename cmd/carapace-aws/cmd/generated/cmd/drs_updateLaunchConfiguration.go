package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_updateLaunchConfigurationCmd = &cobra.Command{
	Use:   "update-launch-configuration",
	Short: "Updates a LaunchConfiguration by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_updateLaunchConfigurationCmd).Standalone()

	drs_updateLaunchConfigurationCmd.Flags().Bool("copy-private-ip", false, "Whether we should copy the Private IP of the Source Server to the Recovery Instance.")
	drs_updateLaunchConfigurationCmd.Flags().Bool("copy-tags", false, "Whether we want to copy the tags of the Source Server to the EC2 machine of the Recovery Instance.")
	drs_updateLaunchConfigurationCmd.Flags().String("launch-disposition", "", "The state of the Recovery Instance in EC2 after the recovery operation.")
	drs_updateLaunchConfigurationCmd.Flags().String("launch-into-instance-properties", "", "Launch into existing instance properties.")
	drs_updateLaunchConfigurationCmd.Flags().String("licensing", "", "The licensing configuration to be used for this launch configuration.")
	drs_updateLaunchConfigurationCmd.Flags().String("name", "", "The name of the launch configuration.")
	drs_updateLaunchConfigurationCmd.Flags().Bool("no-copy-private-ip", false, "Whether we should copy the Private IP of the Source Server to the Recovery Instance.")
	drs_updateLaunchConfigurationCmd.Flags().Bool("no-copy-tags", false, "Whether we want to copy the tags of the Source Server to the EC2 machine of the Recovery Instance.")
	drs_updateLaunchConfigurationCmd.Flags().Bool("no-post-launch-enabled", false, "Whether we want to enable post-launch actions for the Source Server.")
	drs_updateLaunchConfigurationCmd.Flags().Bool("post-launch-enabled", false, "Whether we want to enable post-launch actions for the Source Server.")
	drs_updateLaunchConfigurationCmd.Flags().String("source-server-id", "", "The ID of the Source Server that we want to retrieve a Launch Configuration for.")
	drs_updateLaunchConfigurationCmd.Flags().String("target-instance-type-right-sizing-method", "", "Whether Elastic Disaster Recovery should try to automatically choose the instance type that best matches the OS, CPU, and RAM of your Source Server.")
	drs_updateLaunchConfigurationCmd.Flag("no-copy-private-ip").Hidden = true
	drs_updateLaunchConfigurationCmd.Flag("no-copy-tags").Hidden = true
	drs_updateLaunchConfigurationCmd.Flag("no-post-launch-enabled").Hidden = true
	drs_updateLaunchConfigurationCmd.MarkFlagRequired("source-server-id")
	drsCmd.AddCommand(drs_updateLaunchConfigurationCmd)
}
