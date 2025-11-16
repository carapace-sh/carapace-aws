package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getFindingsPublicationConfigurationCmd = &cobra.Command{
	Use:   "get-findings-publication-configuration",
	Short: "Retrieves the configuration settings for publishing findings to Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getFindingsPublicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getFindingsPublicationConfigurationCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_getFindingsPublicationConfigurationCmd)
}
