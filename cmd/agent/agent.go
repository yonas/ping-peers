package agent

import (
  a "broadcast/pkg/agent"

  "github.com/spf13/cobra"
)

var tags string
var response string

var agentCmd = &cobra.Command{
  Use:   "agent",
  Aliases: []string{"s"},
  Short:  "Listen for broadcast messages",
  //Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    a.Listen(tags, response)
  },
}

func init() {
  agentCmd.Flags().StringVarP(&tags, "tags", "t", "", "Tags")
  agentCmd.Flags().StringVarP(&response, "response", "r", "hey there", "Reply with this message")
}

func AgentCmd() *cobra.Command {
  return agentCmd
}
