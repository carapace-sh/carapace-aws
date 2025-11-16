package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updatePortalCmd = &cobra.Command{
	Use:   "update-portal",
	Short: "Updates an IoT SiteWise Monitor portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updatePortalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_updatePortalCmd).Standalone()

		iotsitewise_updatePortalCmd.Flags().String("alarms", "", "Contains the configuration information of an alarm created in an IoT SiteWise Monitor portal.")
		iotsitewise_updatePortalCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_updatePortalCmd.Flags().String("notification-sender-email", "", "The email address that sends alarm notifications.")
		iotsitewise_updatePortalCmd.Flags().String("portal-contact-email", "", "The Amazon Web Services administrator's contact email address.")
		iotsitewise_updatePortalCmd.Flags().String("portal-description", "", "A new description for the portal.")
		iotsitewise_updatePortalCmd.Flags().String("portal-id", "", "The ID of the portal to update.")
		iotsitewise_updatePortalCmd.Flags().String("portal-logo-image", "", "")
		iotsitewise_updatePortalCmd.Flags().String("portal-name", "", "A new friendly name for the portal.")
		iotsitewise_updatePortalCmd.Flags().String("portal-type", "", "Define the type of portal.")
		iotsitewise_updatePortalCmd.Flags().String("portal-type-configuration", "", "The configuration entry associated with the specific portal type.")
		iotsitewise_updatePortalCmd.Flags().String("role-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a service role that allows the portal's users to access your IoT SiteWise resources on your behalf.")
		iotsitewise_updatePortalCmd.MarkFlagRequired("portal-contact-email")
		iotsitewise_updatePortalCmd.MarkFlagRequired("portal-id")
		iotsitewise_updatePortalCmd.MarkFlagRequired("portal-name")
		iotsitewise_updatePortalCmd.MarkFlagRequired("role-arn")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_updatePortalCmd)
}
