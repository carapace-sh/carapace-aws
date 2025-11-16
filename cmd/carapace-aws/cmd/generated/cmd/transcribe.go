package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Amazon Transcribe offers three main types of batch transcription: **Standard**, **Medical**, and **Call Analytics**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribeCmd).Standalone()

	rootCmd.AddCommand(transcribeCmd)
}
