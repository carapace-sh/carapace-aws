package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "update-launch-configuration-template",
	Short: "Updates an existing Launch Configuration Template by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateLaunchConfigurationTemplateCmd).Standalone()

	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("associate-public-ip-address", false, "Associate public Ip address.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("boot-mode", "", "Launch configuration template boot mode.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("copy-private-ip", false, "Copy private Ip.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("copy-tags", false, "Copy tags.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("large-volume-conf", "", "Large volume config.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("launch-configuration-template-id", "", "Launch Configuration Template ID.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("launch-disposition", "", "Launch disposition.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("licensing", "", "")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("map-auto-tagging-mpe-id", "", "Launch configuration template map auto tagging MPE ID.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-associate-public-ip-address", false, "Associate public Ip address.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-private-ip", false, "Copy private Ip.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-tags", false, "Copy tags.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("post-launch-actions", "", "Post Launch Action to execute on the Test or Cutover instance.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("small-volume-conf", "", "Small volume config.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("small-volume-max-size", "", "Small volume maximum size.")
	mgn_updateLaunchConfigurationTemplateCmd.Flags().String("target-instance-type-right-sizing-method", "", "Target instance type right-sizing method.")
	mgn_updateLaunchConfigurationTemplateCmd.MarkFlagRequired("launch-configuration-template-id")
	mgn_updateLaunchConfigurationTemplateCmd.Flag("no-associate-public-ip-address").Hidden = true
	mgn_updateLaunchConfigurationTemplateCmd.Flag("no-copy-private-ip").Hidden = true
	mgn_updateLaunchConfigurationTemplateCmd.Flag("no-copy-tags").Hidden = true
	mgn_updateLaunchConfigurationTemplateCmd.Flag("no-enable-map-auto-tagging").Hidden = true
	mgnCmd.AddCommand(mgn_updateLaunchConfigurationTemplateCmd)
}
