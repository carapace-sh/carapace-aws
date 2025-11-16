package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_createLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "create-launch-configuration-template",
	Short: "Creates a new Launch Configuration Template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_createLaunchConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_createLaunchConfigurationTemplateCmd).Standalone()

		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("copy-private-ip", false, "Copy private IP.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("copy-tags", false, "Copy tags.")
		drs_createLaunchConfigurationTemplateCmd.Flags().String("export-bucket-arn", "", "S3 bucket ARN to export Source Network templates.")
		drs_createLaunchConfigurationTemplateCmd.Flags().String("launch-disposition", "", "Launch disposition.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("launch-into-source-instance", false, "DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.")
		drs_createLaunchConfigurationTemplateCmd.Flags().String("licensing", "", "Licensing.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-private-ip", false, "Copy private IP.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-tags", false, "Copy tags.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("no-launch-into-source-instance", false, "DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("no-post-launch-enabled", false, "Whether we want to activate post-launch actions.")
		drs_createLaunchConfigurationTemplateCmd.Flags().Bool("post-launch-enabled", false, "Whether we want to activate post-launch actions.")
		drs_createLaunchConfigurationTemplateCmd.Flags().String("tags", "", "Request to associate tags during creation of a Launch Configuration Template.")
		drs_createLaunchConfigurationTemplateCmd.Flags().String("target-instance-type-right-sizing-method", "", "Target instance type right-sizing method.")
		drs_createLaunchConfigurationTemplateCmd.Flag("no-copy-private-ip").Hidden = true
		drs_createLaunchConfigurationTemplateCmd.Flag("no-copy-tags").Hidden = true
		drs_createLaunchConfigurationTemplateCmd.Flag("no-launch-into-source-instance").Hidden = true
		drs_createLaunchConfigurationTemplateCmd.Flag("no-post-launch-enabled").Hidden = true
	})
	drsCmd.AddCommand(drs_createLaunchConfigurationTemplateCmd)
}
