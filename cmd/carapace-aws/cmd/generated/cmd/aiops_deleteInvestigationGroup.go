package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_deleteInvestigationGroupCmd = &cobra.Command{
	Use:   "delete-investigation-group",
	Short: "Deletes the specified investigation group from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_deleteInvestigationGroupCmd).Standalone()

	aiops_deleteInvestigationGroupCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to delete.")
	aiops_deleteInvestigationGroupCmd.MarkFlagRequired("identifier")
	aiopsCmd.AddCommand(aiops_deleteInvestigationGroupCmd)
}
