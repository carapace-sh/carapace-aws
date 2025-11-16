package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_associateAlias2020_05_31Cmd = &cobra.Command{
	Use:   "associate-alias2020_05_31",
	Short: "The `AssociateAlias` API operation only supports standard distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_associateAlias2020_05_31Cmd).Standalone()

	cloudfront_associateAlias2020_05_31Cmd.Flags().String("alias", "", "The alias (also known as a CNAME) to add to the target standard distribution.")
	cloudfront_associateAlias2020_05_31Cmd.Flags().String("target-distribution-id", "", "The ID of the standard distribution that you're associating the alias with.")
	cloudfront_associateAlias2020_05_31Cmd.MarkFlagRequired("alias")
	cloudfront_associateAlias2020_05_31Cmd.MarkFlagRequired("target-distribution-id")
	cloudfrontCmd.AddCommand(cloudfront_associateAlias2020_05_31Cmd)
}
