package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the tag from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_untagResourceCmd).Standalone()

		timestreamInfluxdb_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the tagged resource.")
		timestreamInfluxdb_untagResourceCmd.Flags().String("tag-keys", "", "The keys used to identify the tags.")
		timestreamInfluxdb_untagResourceCmd.MarkFlagRequired("resource-arn")
		timestreamInfluxdb_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_untagResourceCmd)
}
