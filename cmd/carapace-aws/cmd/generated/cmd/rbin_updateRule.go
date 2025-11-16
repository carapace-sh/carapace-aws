package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "Updates an existing Recycle Bin retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_updateRuleCmd).Standalone()

	rbin_updateRuleCmd.Flags().String("description", "", "The retention rule description.")
	rbin_updateRuleCmd.Flags().String("exclude-resource-tags", "", "\\[Region-level retention rules only] Specifies the exclusion tags to use to identify resources that are to be excluded, or ignored, by a Region-level retention rule.")
	rbin_updateRuleCmd.Flags().String("identifier", "", "The unique ID of the retention rule.")
	rbin_updateRuleCmd.Flags().String("resource-tags", "", "\\[Tag-level retention rules only] Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule.")
	rbin_updateRuleCmd.Flags().String("resource-type", "", "This parameter is currently not supported.")
	rbin_updateRuleCmd.Flags().String("retention-period", "", "Information about the retention period for which the retention rule is to retain resources.")
	rbin_updateRuleCmd.MarkFlagRequired("identifier")
	rbinCmd.AddCommand(rbin_updateRuleCmd)
}
