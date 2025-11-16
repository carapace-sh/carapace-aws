package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createAppBlockBuilderCmd = &cobra.Command{
	Use:   "create-app-block-builder",
	Short: "Creates an app block builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createAppBlockBuilderCmd).Standalone()

	appstream_createAppBlockBuilderCmd.Flags().String("access-endpoints", "", "The list of interface VPC endpoint (interface endpoint) objects.")
	appstream_createAppBlockBuilderCmd.Flags().String("description", "", "The description of the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("display-name", "", "The display name of the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("enable-default-internet-access", "", "Enables or disables default internet access for the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to apply to the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("instance-type", "", "The instance type to use when launching the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("name", "", "The unique name for the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("platform", "", "The platform of the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("tags", "", "The tags to associate with the app block builder.")
	appstream_createAppBlockBuilderCmd.Flags().String("vpc-config", "", "The VPC configuration for the app block builder.")
	appstream_createAppBlockBuilderCmd.MarkFlagRequired("instance-type")
	appstream_createAppBlockBuilderCmd.MarkFlagRequired("name")
	appstream_createAppBlockBuilderCmd.MarkFlagRequired("platform")
	appstream_createAppBlockBuilderCmd.MarkFlagRequired("vpc-config")
	appstreamCmd.AddCommand(appstream_createAppBlockBuilderCmd)
}
