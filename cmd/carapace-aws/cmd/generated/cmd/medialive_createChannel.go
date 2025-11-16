package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a new channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createChannelCmd).Standalone()

	medialive_createChannelCmd.Flags().String("anywhere-settings", "", "The Elemental Anywhere settings for this channel.")
	medialive_createChannelCmd.Flags().String("cdi-input-specification", "", "Specification of CDI inputs for this channel")
	medialive_createChannelCmd.Flags().String("channel-class", "", "The class for this channel.")
	medialive_createChannelCmd.Flags().String("channel-engine-version", "", "The desired engine version for this channel.")
	medialive_createChannelCmd.Flags().String("destinations", "", "")
	medialive_createChannelCmd.Flags().String("dry-run", "", "")
	medialive_createChannelCmd.Flags().String("encoder-settings", "", "")
	medialive_createChannelCmd.Flags().String("input-attachments", "", "List of input attachments for channel.")
	medialive_createChannelCmd.Flags().String("input-specification", "", "Specification of network and file inputs for this channel")
	medialive_createChannelCmd.Flags().String("log-level", "", "The log level to write to CloudWatch Logs.")
	medialive_createChannelCmd.Flags().String("maintenance", "", "Maintenance settings for this channel.")
	medialive_createChannelCmd.Flags().String("name", "", "Name of channel.")
	medialive_createChannelCmd.Flags().String("request-id", "", "Unique request ID to be specified.")
	medialive_createChannelCmd.Flags().String("reserved", "", "Deprecated field that's only usable by whitelisted customers.")
	medialive_createChannelCmd.Flags().String("role-arn", "", "An optional Amazon Resource Name (ARN) of the role to assume when running the Channel.")
	medialive_createChannelCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_createChannelCmd.Flags().String("vpc", "", "Settings for the VPC outputs")
	medialiveCmd.AddCommand(medialive_createChannelCmd)
}
