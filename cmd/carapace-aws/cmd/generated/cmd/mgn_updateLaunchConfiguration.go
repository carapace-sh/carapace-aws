package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateLaunchConfigurationCmd = &cobra.Command{
	Use:   "update-launch-configuration",
	Short: "Updates multiple LaunchConfigurations by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateLaunchConfigurationCmd).Standalone()

	mgn_updateLaunchConfigurationCmd.Flags().String("account-id", "", "Update Launch configuration Account ID.")
	mgn_updateLaunchConfigurationCmd.Flags().String("boot-mode", "", "Update Launch configuration boot mode request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("copy-private-ip", false, "Update Launch configuration copy Private IP request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("copy-tags", false, "Update Launch configuration copy Tags request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_updateLaunchConfigurationCmd.Flags().String("launch-disposition", "", "Update Launch configuration launch disposition request.")
	mgn_updateLaunchConfigurationCmd.Flags().String("licensing", "", "Update Launch configuration licensing request.")
	mgn_updateLaunchConfigurationCmd.Flags().String("map-auto-tagging-mpe-id", "", "Launch configuration map auto tagging MPE ID.")
	mgn_updateLaunchConfigurationCmd.Flags().String("name", "", "Update Launch configuration name request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("no-copy-private-ip", false, "Update Launch configuration copy Private IP request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("no-copy-tags", false, "Update Launch configuration copy Tags request.")
	mgn_updateLaunchConfigurationCmd.Flags().Bool("no-enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_updateLaunchConfigurationCmd.Flags().String("post-launch-actions", "", "")
	mgn_updateLaunchConfigurationCmd.Flags().String("source-server-id", "", "Update Launch configuration by Source Server ID request.")
	mgn_updateLaunchConfigurationCmd.Flags().String("target-instance-type-right-sizing-method", "", "Update Launch configuration Target instance right sizing request.")
	mgn_updateLaunchConfigurationCmd.Flag("no-copy-private-ip").Hidden = true
	mgn_updateLaunchConfigurationCmd.Flag("no-copy-tags").Hidden = true
	mgn_updateLaunchConfigurationCmd.Flag("no-enable-map-auto-tagging").Hidden = true
	mgn_updateLaunchConfigurationCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_updateLaunchConfigurationCmd)
}
