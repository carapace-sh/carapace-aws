package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listKeyGroups2020_05_31Cmd = &cobra.Command{
	Use:   "list-key-groups2020_05_31",
	Short: "Gets a list of key groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listKeyGroups2020_05_31Cmd).Standalone()

	cloudfront_listKeyGroups2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of key groups.")
	cloudfront_listKeyGroups2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of key groups that you want in the response.")
	cloudfrontCmd.AddCommand(cloudfront_listKeyGroups2020_05_31Cmd)
}
