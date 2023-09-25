package send

import (
  "fmt"
  "net"
  "strings"

  "broadcast/pkg/agent"
  log "broadcast/pkg/logging"

  "github.com/spf13/cobra"
  "github.com/vmihailenco/msgpack/v5"
)

var agentTags string
var response string

func init() {
  agentCmd.Flags().StringVarP(&agentTags, "tags", "t", "", "Tags")
  agentCmd.Flags().StringVarP(&response, "response", "r", "hey there", "Reply with this message")

  rootCmd.AddCommand(agentCmd)
}

var agentCmd = &cobra.Command{
  Use:   "agent",
  Aliases: []string{"s"},
  Short:  "Listen for broadcast messages",
  //Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    agent.listen(agentTags, response)
  },
}
