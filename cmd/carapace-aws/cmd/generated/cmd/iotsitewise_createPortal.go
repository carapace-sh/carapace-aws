package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createPortalCmd = &cobra.Command{
	Use:   "create-portal",
	Short: "Creates a portal, which can contain projects and dashboards.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createPortalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_createPortalCmd).Standalone()

		iotsitewise_createPortalCmd.Flags().String("alarms", "", "Contains the configuration information of an alarm created in an IoT SiteWise Monitor portal.")
		iotsitewise_createPortalCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_createPortalCmd.Flags().String("notification-sender-email", "", "The email address that sends alarm notifications.")
		iotsitewise_createPortalCmd.Flags().String("portal-auth-mode", "", "The service to use to authenticate users to the portal.")
		iotsitewise_createPortalCmd.Flags().String("portal-contact-email", "", "The Amazon Web Services administrator's contact email address.")
		iotsitewise_createPortalCmd.Flags().String("portal-description", "", "A description for the portal.")
		iotsitewise_createPortalCmd.Flags().String("portal-logo-image-file", "", "A logo image to display in the portal.")
		iotsitewise_createPortalCmd.Flags().String("portal-name", "", "A friendly name for the portal.")
		iotsitewise_createPortalCmd.Flags().String("portal-type", "", "Define the type of portal.")
		iotsitewise_createPortalCmd.Flags().String("portal-type-configuration", "", "The configuration entry associated with the specific portal type.")
		iotsitewise_createPortalCmd.Flags().String("role-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your IoT SiteWise resources on your behalf.")
		iotsitewise_createPortalCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the portal.")
		iotsitewise_createPortalCmd.MarkFlagRequired("portal-contact-email")
		iotsitewise_createPortalCmd.MarkFlagRequired("portal-name")
		iotsitewise_createPortalCmd.MarkFlagRequired("role-arn")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_createPortalCmd)
}
