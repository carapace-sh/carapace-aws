package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd = &cobra.Command{
	Use:   "delete-application-vpc-configuration",
	Short: "Removes a VPC configuration from a Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd).Standalone()

	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.Flags().String("application-name", "", "The name of an existing application.")
	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.Flags().String("conditional-token", "", "A value you use to implement strong concurrency for application updates.")
	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.Flags().String("current-application-version-id", "", "The current application version ID.")
	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.Flags().String("vpc-configuration-id", "", "The ID of the VPC configuration to delete.")
	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd.MarkFlagRequired("vpc-configuration-id")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationVpcConfigurationCmd)
}
