package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getMasterAccountCmd = &cobra.Command{
	Use:   "get-master-account",
	Short: "(Deprecated) Retrieves information about the Amazon Macie administrator account for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getMasterAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getMasterAccountCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_getMasterAccountCmd)
}
