package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createImageBuilderCmd = &cobra.Command{
	Use:   "create-image-builder",
	Short: "Creates an image builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createImageBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createImageBuilderCmd).Standalone()

		appstream_createImageBuilderCmd.Flags().String("access-endpoints", "", "The list of interface VPC endpoint (interface endpoint) objects.")
		appstream_createImageBuilderCmd.Flags().String("appstream-agent-version", "", "The version of the AppStream 2.0 agent to use for this image builder.")
		appstream_createImageBuilderCmd.Flags().String("description", "", "The description to display.")
		appstream_createImageBuilderCmd.Flags().String("display-name", "", "The image builder name to display.")
		appstream_createImageBuilderCmd.Flags().String("domain-join-info", "", "The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.")
		appstream_createImageBuilderCmd.Flags().String("enable-default-internet-access", "", "Enables or disables default internet access for the image builder.")
		appstream_createImageBuilderCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to apply to the image builder.")
		appstream_createImageBuilderCmd.Flags().String("image-arn", "", "The ARN of the public, private, or shared image to use.")
		appstream_createImageBuilderCmd.Flags().String("image-name", "", "The name of the image used to create the image builder.")
		appstream_createImageBuilderCmd.Flags().String("instance-type", "", "The instance type to use when launching the image builder.")
		appstream_createImageBuilderCmd.Flags().String("name", "", "A unique name for the image builder.")
		appstream_createImageBuilderCmd.Flags().String("softwares-to-install", "", "The list of license included applications to install on the image builder during creation.")
		appstream_createImageBuilderCmd.Flags().String("softwares-to-uninstall", "", "The list of license included applications to uninstall from the image builder during creation.")
		appstream_createImageBuilderCmd.Flags().String("tags", "", "The tags to associate with the image builder.")
		appstream_createImageBuilderCmd.Flags().String("vpc-config", "", "The VPC configuration for the image builder.")
		appstream_createImageBuilderCmd.MarkFlagRequired("instance-type")
		appstream_createImageBuilderCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_createImageBuilderCmd)
}
