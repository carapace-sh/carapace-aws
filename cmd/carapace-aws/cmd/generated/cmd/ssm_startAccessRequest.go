package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startAccessRequestCmd = &cobra.Command{
	Use:   "start-access-request",
	Short: "Starts the workflow for just-in-time node access sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startAccessRequestCmd).Standalone()

	ssm_startAccessRequestCmd.Flags().String("reason", "", "A brief description explaining why you are requesting access to the node.")
	ssm_startAccessRequestCmd.Flags().String("tags", "", "Key-value pairs of metadata you want to assign to the access request.")
	ssm_startAccessRequestCmd.Flags().String("targets", "", "The node you are requesting access to.")
	ssm_startAccessRequestCmd.MarkFlagRequired("reason")
	ssm_startAccessRequestCmd.MarkFlagRequired("targets")
	ssmCmd.AddCommand(ssm_startAccessRequestCmd)
}
