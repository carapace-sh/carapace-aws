package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getCurrentUserCmd = &cobra.Command{
	Use:   "get-current-user",
	Short: "Retrieves details of the current user for whom the authentication token was generated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getCurrentUserCmd).Standalone()

	workdocs_getCurrentUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_getCurrentUserCmd.MarkFlagRequired("authentication-token")
	workdocsCmd.AddCommand(workdocs_getCurrentUserCmd)
}
