package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeEventTopicsCmd = &cobra.Command{
	Use:   "describe-event-topics",
	Short: "Obtains information about which Amazon SNS topics receive status messages from the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeEventTopicsCmd).Standalone()

	ds_describeEventTopicsCmd.Flags().String("directory-id", "", "The Directory ID for which to get the list of associated Amazon SNS topics.")
	ds_describeEventTopicsCmd.Flags().String("topic-names", "", "A list of Amazon SNS topic names for which to obtain the information.")
	dsCmd.AddCommand(ds_describeEventTopicsCmd)
}
