package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_enableMacieCmd = &cobra.Command{
	Use:   "enable-macie",
	Short: "Enables Amazon Macie and specifies the configuration settings for a Macie account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_enableMacieCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_enableMacieCmd).Standalone()

		macie2_enableMacieCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		macie2_enableMacieCmd.Flags().String("finding-publishing-frequency", "", "Specifies how often to publish updates to policy findings for the account.")
		macie2_enableMacieCmd.Flags().String("status", "", "Specifies the new status for the account.")
	})
	macie2Cmd.AddCommand(macie2_enableMacieCmd)
}
