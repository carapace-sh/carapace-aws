package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createFaceLivenessSessionCmd = &cobra.Command{
	Use:   "create-face-liveness-session",
	Short: "This API operation initiates a Face Liveness session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createFaceLivenessSessionCmd).Standalone()

	rekognition_createFaceLivenessSessionCmd.Flags().String("client-request-token", "", "Idempotent token is used to recognize the Face Liveness request.")
	rekognition_createFaceLivenessSessionCmd.Flags().String("kms-key-id", "", "The identifier for your AWS Key Management Service key (AWS KMS key).")
	rekognition_createFaceLivenessSessionCmd.Flags().String("settings", "", "A session settings object.")
	rekognitionCmd.AddCommand(rekognition_createFaceLivenessSessionCmd)
}
