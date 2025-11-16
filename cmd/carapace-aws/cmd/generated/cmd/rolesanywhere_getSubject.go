package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_getSubjectCmd = &cobra.Command{
	Use:   "get-subject",
	Short: "Gets a *subject*, which associates a certificate identity with authentication attempts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_getSubjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_getSubjectCmd).Standalone()

		rolesanywhere_getSubjectCmd.Flags().String("subject-id", "", "The unique identifier of the subject.")
		rolesanywhere_getSubjectCmd.MarkFlagRequired("subject-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_getSubjectCmd)
}
