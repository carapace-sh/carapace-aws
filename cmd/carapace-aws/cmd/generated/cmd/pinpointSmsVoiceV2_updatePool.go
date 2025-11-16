package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_updatePoolCmd = &cobra.Command{
	Use:   "update-pool",
	Short: "Updates the configuration of an existing pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_updatePoolCmd).Standalone()

	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("deletion-protection-enabled", false, "When set to true the pool can't be deleted.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("no-deletion-protection-enabled", false, "When set to true the pool can't be deleted.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("no-self-managed-opt-outs-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("no-shared-routes-enabled", false, "Indicates whether shared routes are enabled for the pool.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("no-two-way-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().String("opt-out-list-name", "", "The OptOutList to associate with the pool.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().String("pool-id", "", "The unique identifier of the pool to update.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("self-managed-opt-outs-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("shared-routes-enabled", false, "Indicates whether shared routes are enabled for the pool.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().String("two-way-channel-arn", "", "The Amazon Resource Name (ARN) of the two way channel.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().String("two-way-channel-role", "", "An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.")
	pinpointSmsVoiceV2_updatePoolCmd.Flags().Bool("two-way-enabled", false, "By default this is set to false.")
	pinpointSmsVoiceV2_updatePoolCmd.Flag("no-deletion-protection-enabled").Hidden = true
	pinpointSmsVoiceV2_updatePoolCmd.Flag("no-self-managed-opt-outs-enabled").Hidden = true
	pinpointSmsVoiceV2_updatePoolCmd.Flag("no-shared-routes-enabled").Hidden = true
	pinpointSmsVoiceV2_updatePoolCmd.Flag("no-two-way-enabled").Hidden = true
	pinpointSmsVoiceV2_updatePoolCmd.MarkFlagRequired("pool-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_updatePoolCmd)
}
