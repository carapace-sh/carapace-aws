package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getMacieSessionCmd = &cobra.Command{
	Use:   "get-macie-session",
	Short: "Retrieves the status and configuration settings for an Amazon Macie account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getMacieSessionCmd).Standalone()

	macie2Cmd.AddCommand(macie2_getMacieSessionCmd)
}
