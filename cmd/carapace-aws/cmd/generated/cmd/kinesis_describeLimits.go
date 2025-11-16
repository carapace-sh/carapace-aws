package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_describeLimitsCmd = &cobra.Command{
	Use:   "describe-limits",
	Short: "Describes the shard limits and usage for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_describeLimitsCmd).Standalone()

	kinesisCmd.AddCommand(kinesis_describeLimitsCmd)
}
