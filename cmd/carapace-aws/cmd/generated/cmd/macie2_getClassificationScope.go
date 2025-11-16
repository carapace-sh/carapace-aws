package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getClassificationScopeCmd = &cobra.Command{
	Use:   "get-classification-scope",
	Short: "Retrieves the classification scope settings for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getClassificationScopeCmd).Standalone()

	macie2_getClassificationScopeCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_getClassificationScopeCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_getClassificationScopeCmd)
}
