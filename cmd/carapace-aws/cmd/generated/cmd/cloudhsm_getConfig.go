package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_getConfigCmd).Standalone()

	cloudhsm_getConfigCmd.Flags().String("client-arn", "", "The ARN of the client.")
	cloudhsm_getConfigCmd.Flags().String("client-version", "", "The client version.")
	cloudhsm_getConfigCmd.Flags().String("hapg-list", "", "A list of ARNs that identify the high-availability partition groups that are associated with the client.")
	cloudhsm_getConfigCmd.MarkFlagRequired("client-arn")
	cloudhsm_getConfigCmd.MarkFlagRequired("client-version")
	cloudhsm_getConfigCmd.MarkFlagRequired("hapg-list")
	cloudhsmCmd.AddCommand(cloudhsm_getConfigCmd)
}
