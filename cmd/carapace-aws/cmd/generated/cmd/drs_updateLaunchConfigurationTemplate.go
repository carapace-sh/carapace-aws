package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_updateLaunchConfigurationTemplateCmd = &cobra.Command{
	Use:   "update-launch-configuration-template",
	Short: "Updates an existing Launch Configuration Template by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_updateLaunchConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_updateLaunchConfigurationTemplateCmd).Standalone()

		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("copy-private-ip", false, "Copy private IP.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("copy-tags", false, "Copy tags.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().String("export-bucket-arn", "", "S3 bucket ARN to export Source Network templates.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().String("launch-configuration-template-id", "", "Launch Configuration Template ID.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().String("launch-disposition", "", "Launch disposition.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("launch-into-source-instance", false, "DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().String("licensing", "", "Licensing.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-private-ip", false, "Copy private IP.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-copy-tags", false, "Copy tags.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-launch-into-source-instance", false, "DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("no-post-launch-enabled", false, "Whether we want to activate post-launch actions.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().Bool("post-launch-enabled", false, "Whether we want to activate post-launch actions.")
		drs_updateLaunchConfigurationTemplateCmd.Flags().String("target-instance-type-right-sizing-method", "", "Target instance type right-sizing method.")
		drs_updateLaunchConfigurationTemplateCmd.MarkFlagRequired("launch-configuration-template-id")
		drs_updateLaunchConfigurationTemplateCmd.Flag("no-copy-private-ip").Hidden = true
		drs_updateLaunchConfigurationTemplateCmd.Flag("no-copy-tags").Hidden = true
		drs_updateLaunchConfigurationTemplateCmd.Flag("no-launch-into-source-instance").Hidden = true
		drs_updateLaunchConfigurationTemplateCmd.Flag("no-post-launch-enabled").Hidden = true
	})
	drsCmd.AddCommand(drs_updateLaunchConfigurationTemplateCmd)
}
