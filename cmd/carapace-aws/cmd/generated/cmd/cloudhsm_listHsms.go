package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_listHsmsCmd = &cobra.Command{
	Use:   "list-hsms",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_listHsmsCmd).Standalone()

	cloudhsm_listHsmsCmd.Flags().String("next-token", "", "The `NextToken` value from a previous call to `ListHsms`.")
	cloudhsmCmd.AddCommand(cloudhsm_listHsmsCmd)
}
