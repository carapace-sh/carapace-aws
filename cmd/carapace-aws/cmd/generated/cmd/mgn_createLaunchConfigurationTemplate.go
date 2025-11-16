package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_createLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "create-launch-configuration-template",
	Short: "Creates a new Launch Configuration Template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_createLaunchConfigurationTemplateCmd).Standalone()

	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("associate-public-ip-address", false, "Associate public Ip address.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("boot-mode", "", "Launch configuration template boot mode.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("copy-private-ip", false, "Copy private Ip.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("copy-tags", false, "Copy tags.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("large-volume-conf", "", "Large volume config.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("launch-disposition", "", "Launch disposition.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("licensing", "", "")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("map-auto-tagging-mpe-id", "", "Launch configuration template map auto tagging MPE ID.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("no-associate-public-ip-address", false, "Associate public Ip address.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-private-ip", false, "Copy private Ip.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-tags", false, "Copy tags.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().Bool("no-enable-map-auto-tagging", false, "Enable map auto tagging.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("post-launch-actions", "", "Launch configuration template post launch actions.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("small-volume-conf", "", "Small volume config.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("small-volume-max-size", "", "Small volume maximum size.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("tags", "", "Request to associate tags during creation of a Launch Configuration Template.")
	mgn_createLaunchConfigurationTemplateCmd.Flags().String("target-instance-type-right-sizing-method", "", "Target instance type right-sizing method.")
	mgn_createLaunchConfigurationTemplateCmd.Flag("no-associate-public-ip-address").Hidden = true
	mgn_createLaunchConfigurationTemplateCmd.Flag("no-copy-private-ip").Hidden = true
	mgn_createLaunchConfigurationTemplateCmd.Flag("no-copy-tags").Hidden = true
	mgn_createLaunchConfigurationTemplateCmd.Flag("no-enable-map-auto-tagging").Hidden = true
	mgnCmd.AddCommand(mgn_createLaunchConfigurationTemplateCmd)
}
