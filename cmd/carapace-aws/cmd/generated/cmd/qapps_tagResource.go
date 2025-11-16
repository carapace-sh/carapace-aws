package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates tags with an Amazon Q Apps resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_tagResourceCmd).Standalone()

	qapps_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
	qapps_tagResourceCmd.Flags().String("tags", "", "The tags to associate with the resource.")
	qapps_tagResourceCmd.MarkFlagRequired("resource-arn")
	qapps_tagResourceCmd.MarkFlagRequired("tags")
	qappsCmd.AddCommand(qapps_tagResourceCmd)
}
