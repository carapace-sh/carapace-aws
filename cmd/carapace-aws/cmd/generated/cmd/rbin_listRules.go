package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "Lists the Recycle Bin retention rules in the Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_listRulesCmd).Standalone()

	rbin_listRulesCmd.Flags().String("exclude-resource-tags", "", "\\[Region-level retention rules only] Information about the exclusion tags used to identify resources that are to be excluded, or ignored, by the retention rule.")
	rbin_listRulesCmd.Flags().String("lock-state", "", "The lock state of the retention rules to list.")
	rbin_listRulesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	rbin_listRulesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	rbin_listRulesCmd.Flags().String("resource-tags", "", "\\[Tag-level retention rules only] Information about the resource tags used to identify resources that are retained by the retention rule.")
	rbin_listRulesCmd.Flags().String("resource-type", "", "The resource type retained by the retention rule.")
	rbin_listRulesCmd.MarkFlagRequired("resource-type")
	rbinCmd.AddCommand(rbin_listRulesCmd)
}
