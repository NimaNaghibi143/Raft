# The Algorithm

Nodes in a cluster conduct elections to pick a leader. Users of the Raft cluster send messages to the leader. The leader passes the message to the followers and waites for a majority to store the message. Once the message is committed (majority consensus has been reached), the message is applied to a state machine the user supplies. Followers learn about the latest committed message forn the leader and apply each new committed message to their local user supplied state machine.

## Components

* **A distributed key-value store:** We need to create a state machine and commands that are sent to the state machine.

* **HTTP API:** We need HTTP endpoints that allow the user tp operate the state mahcine through the Raft cluster.
