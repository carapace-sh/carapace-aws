package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A list of tags applied to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_listTagsForResourceCmd).Standalone()

		timestreamInfluxdb_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the tagged resource.")
		timestreamInfluxdb_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_listTagsForResourceCmd)
}
