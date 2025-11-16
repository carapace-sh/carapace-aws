package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "Creates a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createFleetCmd).Standalone()

	appstream_createFleetCmd.Flags().String("compute-capacity", "", "The desired capacity for the fleet.")
	appstream_createFleetCmd.Flags().String("description", "", "The description to display.")
	appstream_createFleetCmd.Flags().String("disconnect-timeout-in-seconds", "", "The amount of time that a streaming session remains active after users disconnect.")
	appstream_createFleetCmd.Flags().String("display-name", "", "The fleet name to display.")
	appstream_createFleetCmd.Flags().String("domain-join-info", "", "The name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain.")
	appstream_createFleetCmd.Flags().String("enable-default-internet-access", "", "Enables or disables default internet access for the fleet.")
	appstream_createFleetCmd.Flags().String("fleet-type", "", "The fleet type.")
	appstream_createFleetCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to apply to the fleet.")
	appstream_createFleetCmd.Flags().String("idle-disconnect-timeout-in-seconds", "", "The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the `DisconnectTimeoutInSeconds` time interval begins.")
	appstream_createFleetCmd.Flags().String("image-arn", "", "The ARN of the public, private, or shared image to use.")
	appstream_createFleetCmd.Flags().String("image-name", "", "The name of the image used to create the fleet.")
	appstream_createFleetCmd.Flags().String("instance-type", "", "The instance type to use when launching fleet instances.")
	appstream_createFleetCmd.Flags().String("max-concurrent-sessions", "", "The maximum concurrent sessions of the Elastic fleet.")
	appstream_createFleetCmd.Flags().String("max-sessions-per-instance", "", "The maximum number of user sessions on an instance.")
	appstream_createFleetCmd.Flags().String("max-user-duration-in-seconds", "", "The maximum amount of time that a streaming session can remain active, in seconds.")
	appstream_createFleetCmd.Flags().String("name", "", "A unique name for the fleet.")
	appstream_createFleetCmd.Flags().String("platform", "", "The fleet platform.")
	appstream_createFleetCmd.Flags().String("session-script-s3-location", "", "The S3 location of the session scripts configuration zip file.")
	appstream_createFleetCmd.Flags().String("stream-view", "", "The AppStream 2.0 view that is displayed to your users when they stream from the fleet.")
	appstream_createFleetCmd.Flags().String("tags", "", "The tags to associate with the fleet.")
	appstream_createFleetCmd.Flags().String("usb-device-filter-strings", "", "The USB device filter strings that specify which USB devices a user can redirect to the fleet streaming session, when using the Windows native client.")
	appstream_createFleetCmd.Flags().String("vpc-config", "", "The VPC configuration for the fleet.")
	appstream_createFleetCmd.MarkFlagRequired("instance-type")
	appstream_createFleetCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_createFleetCmd)
}
