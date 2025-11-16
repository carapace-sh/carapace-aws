package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointCmd = &cobra.Command{
	Use:   "pinpoint",
	Short: "Doc Engage API - Amazon Pinpoint API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointCmd).Standalone()

	rootCmd.AddCommand(pinpointCmd)
}
