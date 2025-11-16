package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getAdministratorAccountCmd = &cobra.Command{
	Use:   "get-administrator-account",
	Short: "Retrieves information about the Amazon Macie administrator account for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getAdministratorAccountCmd).Standalone()

	macie2Cmd.AddCommand(macie2_getAdministratorAccountCmd)
}
