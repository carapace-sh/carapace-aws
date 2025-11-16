package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteCaseCmd = &cobra.Command{
	Use:   "delete-case",
	Short: "The DeleteCase API permanently deletes a case and all its associated resources from the cases data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteCaseCmd).Standalone()

	connectcases_deleteCaseCmd.Flags().String("case-id", "", "A unique identifier of the case.")
	connectcases_deleteCaseCmd.Flags().String("domain-id", "", "A unique identifier of the Cases domain.")
	connectcases_deleteCaseCmd.MarkFlagRequired("case-id")
	connectcases_deleteCaseCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_deleteCaseCmd)
}
