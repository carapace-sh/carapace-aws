package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getAllowListCmd = &cobra.Command{
	Use:   "get-allow-list",
	Short: "Retrieves the settings and status of an allow list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getAllowListCmd).Standalone()

	macie2_getAllowListCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_getAllowListCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_getAllowListCmd)
}
