package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags are composed of a Key/Value pairs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_tagResourceCmd).Standalone()

	timestreamInfluxdb_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the tagged resource.")
	timestreamInfluxdb_tagResourceCmd.Flags().String("tags", "", "A list of tags used to categorize and track resources.")
	timestreamInfluxdb_tagResourceCmd.MarkFlagRequired("resource-arn")
	timestreamInfluxdb_tagResourceCmd.MarkFlagRequired("tags")
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_tagResourceCmd)
}
