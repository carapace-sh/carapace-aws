package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getSendQuotaCmd = &cobra.Command{
	Use:   "get-send-quota",
	Short: "Provides the sending limits for the Amazon SES account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getSendQuotaCmd).Standalone()

	sesCmd.AddCommand(ses_getSendQuotaCmd)
}
