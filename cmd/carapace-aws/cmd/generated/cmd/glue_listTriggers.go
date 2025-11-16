package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listTriggersCmd = &cobra.Command{
	Use:   "list-triggers",
	Short: "Retrieves the names of all trigger resources in this Amazon Web Services account, or the resources with the specified tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listTriggersCmd).Standalone()

	glue_listTriggersCmd.Flags().String("dependent-job-name", "", "The name of the job for which to retrieve triggers.")
	glue_listTriggersCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
	glue_listTriggersCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
	glue_listTriggersCmd.Flags().String("tags", "", "Specifies to return only these tagged resources.")
	glueCmd.AddCommand(glue_listTriggersCmd)
}
