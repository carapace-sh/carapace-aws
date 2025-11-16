package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_describeConnectionTypeCmd = &cobra.Command{
	Use:   "describe-connection-type",
	Short: "The `DescribeConnectionType` API provides full details of the supported options for a given connection type in Glue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_describeConnectionTypeCmd).Standalone()

	glue_describeConnectionTypeCmd.Flags().String("connection-type", "", "The name of the connection type to be described.")
	glue_describeConnectionTypeCmd.MarkFlagRequired("connection-type")
	glueCmd.AddCommand(glue_describeConnectionTypeCmd)
}
