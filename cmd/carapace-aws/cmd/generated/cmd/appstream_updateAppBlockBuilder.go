package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateAppBlockBuilderCmd = &cobra.Command{
	Use:   "update-app-block-builder",
	Short: "Updates an app block builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateAppBlockBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_updateAppBlockBuilderCmd).Standalone()

		appstream_updateAppBlockBuilderCmd.Flags().String("access-endpoints", "", "The list of interface VPC endpoint (interface endpoint) objects.")
		appstream_updateAppBlockBuilderCmd.Flags().String("attributes-to-delete", "", "The attributes to delete from the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("description", "", "The description of the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("display-name", "", "The display name of the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("enable-default-internet-access", "", "Enables or disables default internet access for the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to apply to the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("instance-type", "", "The instance type to use when launching the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("name", "", "The unique name for the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("platform", "", "The platform of the app block builder.")
		appstream_updateAppBlockBuilderCmd.Flags().String("vpc-config", "", "The VPC configuration for the app block builder.")
		appstream_updateAppBlockBuilderCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_updateAppBlockBuilderCmd)
}
