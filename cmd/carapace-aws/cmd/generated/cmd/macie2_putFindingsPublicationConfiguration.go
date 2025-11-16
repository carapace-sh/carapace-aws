package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_putFindingsPublicationConfigurationCmd = &cobra.Command{
	Use:   "put-findings-publication-configuration",
	Short: "Updates the configuration settings for publishing findings to Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_putFindingsPublicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_putFindingsPublicationConfigurationCmd).Standalone()

		macie2_putFindingsPublicationConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		macie2_putFindingsPublicationConfigurationCmd.Flags().String("security-hub-configuration", "", "The configuration settings that determine which findings to publish to Security Hub.")
	})
	macie2Cmd.AddCommand(macie2_putFindingsPublicationConfigurationCmd)
}
