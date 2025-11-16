package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_getInvestigationGroupCmd = &cobra.Command{
	Use:   "get-investigation-group",
	Short: "Returns the configuration information for the specified investigation group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_getInvestigationGroupCmd).Standalone()

	aiops_getInvestigationGroupCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to view.")
	aiops_getInvestigationGroupCmd.MarkFlagRequired("identifier")
	aiopsCmd.AddCommand(aiops_getInvestigationGroupCmd)
}
