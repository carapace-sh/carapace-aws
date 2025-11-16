package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createFindingsFilterCmd = &cobra.Command{
	Use:   "create-findings-filter",
	Short: "Creates and defines the criteria and other settings for a findings filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createFindingsFilterCmd).Standalone()

	macie2_createFindingsFilterCmd.Flags().String("action", "", "The action to perform on findings that match the filter criteria (findingCriteria).")
	macie2_createFindingsFilterCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
	macie2_createFindingsFilterCmd.Flags().String("description", "", "A custom description of the filter.")
	macie2_createFindingsFilterCmd.Flags().String("finding-criteria", "", "The criteria to use to filter findings.")
	macie2_createFindingsFilterCmd.Flags().String("name", "", "A custom name for the filter.")
	macie2_createFindingsFilterCmd.Flags().String("position", "", "The position of the filter in the list of saved filters on the Amazon Macie console.")
	macie2_createFindingsFilterCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the filter.")
	macie2_createFindingsFilterCmd.MarkFlagRequired("action")
	macie2_createFindingsFilterCmd.MarkFlagRequired("finding-criteria")
	macie2_createFindingsFilterCmd.MarkFlagRequired("name")
	macie2Cmd.AddCommand(macie2_createFindingsFilterCmd)
}
