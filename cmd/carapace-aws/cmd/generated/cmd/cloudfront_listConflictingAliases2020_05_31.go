package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listConflictingAliases2020_05_31Cmd = &cobra.Command{
	Use:   "list-conflicting-aliases2020_05_31",
	Short: "The `ListConflictingAliases` API operation only supports standard distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listConflictingAliases2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listConflictingAliases2020_05_31Cmd).Standalone()

		cloudfront_listConflictingAliases2020_05_31Cmd.Flags().String("alias", "", "The alias (also called a CNAME) to search for conflicting aliases.")
		cloudfront_listConflictingAliases2020_05_31Cmd.Flags().String("distribution-id", "", "The ID of a standard distribution in your account that has an attached TLS certificate that includes the provided alias.")
		cloudfront_listConflictingAliases2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in the list of conflicting aliases.")
		cloudfront_listConflictingAliases2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of conflicting aliases that you want in the response.")
		cloudfront_listConflictingAliases2020_05_31Cmd.MarkFlagRequired("alias")
		cloudfront_listConflictingAliases2020_05_31Cmd.MarkFlagRequired("distribution-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listConflictingAliases2020_05_31Cmd)
}
