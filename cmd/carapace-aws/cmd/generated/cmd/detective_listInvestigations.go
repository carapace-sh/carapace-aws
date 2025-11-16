package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listInvestigationsCmd = &cobra.Command{
	Use:   "list-investigations",
	Short: "Detective investigations lets you investigate IAM users and IAM roles using indicators of compromise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listInvestigationsCmd).Standalone()

	detective_listInvestigationsCmd.Flags().String("filter-criteria", "", "Filters the investigation results based on a criteria.")
	detective_listInvestigationsCmd.Flags().String("graph-arn", "", "The Amazon Resource Name (ARN) of the behavior graph.")
	detective_listInvestigationsCmd.Flags().String("max-results", "", "Lists the maximum number of investigations in a page.")
	detective_listInvestigationsCmd.Flags().String("next-token", "", "Lists if there are more results available.")
	detective_listInvestigationsCmd.Flags().String("sort-criteria", "", "Sorts the investigation results based on a criteria.")
	detective_listInvestigationsCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_listInvestigationsCmd)
}
