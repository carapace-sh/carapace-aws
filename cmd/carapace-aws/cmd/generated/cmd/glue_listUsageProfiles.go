package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listUsageProfilesCmd = &cobra.Command{
	Use:   "list-usage-profiles",
	Short: "List all the Glue usage profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listUsageProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listUsageProfilesCmd).Standalone()

		glue_listUsageProfilesCmd.Flags().String("max-results", "", "The maximum number of usage profiles to return in a single response.")
		glue_listUsageProfilesCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_listUsageProfilesCmd)
}
