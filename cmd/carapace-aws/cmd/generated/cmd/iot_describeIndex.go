package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeIndexCmd = &cobra.Command{
	Use:   "describe-index",
	Short: "Describes a search index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeIndexCmd).Standalone()

		iot_describeIndexCmd.Flags().String("index-name", "", "The index name.")
		iot_describeIndexCmd.MarkFlagRequired("index-name")
	})
	iotCmd.AddCommand(iot_describeIndexCmd)
}
