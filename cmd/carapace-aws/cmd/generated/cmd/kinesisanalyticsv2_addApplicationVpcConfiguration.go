package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationVpcConfigurationCmd = &cobra.Command{
	Use:   "add-application-vpc-configuration",
	Short: "Adds a Virtual Private Cloud (VPC) configuration to the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationVpcConfigurationCmd).Standalone()

	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.Flags().String("application-name", "", "The name of an existing application.")
	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.Flags().String("conditional-token", "", "A value you use to implement strong concurrency for application updates.")
	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.Flags().String("current-application-version-id", "", "The version of the application to which you want to add the VPC configuration.")
	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.Flags().String("vpc-configuration", "", "Description of the VPC to add to the application.")
	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_addApplicationVpcConfigurationCmd.MarkFlagRequired("vpc-configuration")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationVpcConfigurationCmd)
}
