package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_resumeReplicationCmd = &cobra.Command{
	Use:   "resume-replication",
	Short: "Resume Replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_resumeReplicationCmd).Standalone()

	mgn_resumeReplicationCmd.Flags().String("account-id", "", "Resume Replication Request account ID.")
	mgn_resumeReplicationCmd.Flags().String("source-server-id", "", "Resume Replication Request source server ID.")
	mgn_resumeReplicationCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_resumeReplicationCmd)
}
