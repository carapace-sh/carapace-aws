package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetTriggersCmd = &cobra.Command{
	Use:   "batch-get-triggers",
	Short: "Returns a list of resource metadata for a given list of trigger names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetTriggersCmd).Standalone()

	glue_batchGetTriggersCmd.Flags().String("trigger-names", "", "A list of trigger names, which may be the names returned from the `ListTriggers` operation.")
	glue_batchGetTriggersCmd.MarkFlagRequired("trigger-names")
	glueCmd.AddCommand(glue_batchGetTriggersCmd)
}
