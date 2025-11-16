package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_deleteLunaClientCmd = &cobra.Command{
	Use:   "delete-luna-client",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_deleteLunaClientCmd).Standalone()

	cloudhsm_deleteLunaClientCmd.Flags().String("client-arn", "", "The ARN of the client to delete.")
	cloudhsm_deleteLunaClientCmd.MarkFlagRequired("client-arn")
	cloudhsmCmd.AddCommand(cloudhsm_deleteLunaClientCmd)
}
