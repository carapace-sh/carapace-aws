package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deletePoolCmd = &cobra.Command{
	Use:   "delete-pool",
	Short: "Deletes an existing pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deletePoolCmd).Standalone()

	pinpointSmsVoiceV2_deletePoolCmd.Flags().String("pool-id", "", "The PoolId or PoolArn of the pool to delete.")
	pinpointSmsVoiceV2_deletePoolCmd.MarkFlagRequired("pool-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deletePoolCmd)
}
