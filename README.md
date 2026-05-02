# The Algorithm

Nodes in a cluster conduct elections to pick a leader. Users of the Raft cluster send messages to the leader. The leader passes the message to the followers and waites for a majority to store the message. Once the message is committed (majority consensus has been reached), the message is applied to a state machine the user supplies. Followers learn about the latest committed message forn the leader and apply each new committed message to their local user supplied state machine.

## Components

* **A distributed key-value store:** We need to create a state machine and commands that are sent to the state machine.

* **HTTP API:** We need HTTP endpoints that allow the user tp operate the state mahcine through the Raft cluster.

**Raft Server** based on the raft papar:

### State

#### Persistent state on all servers

* currentTerm: Latest term server has seen.
* votedFor: candidateId that recieved vote in current term(or null if none).
* log[]: log entries; each entry contains command for state machine, and term when entry was recieved by leader.(first index is 1)
  
#### Volatile state on all servers

* commitIndex: index of highest log entry known to be committed.
* lastApplied: index of the highest log entry applied to state machine.

#### Volatile state on leaders

* nextIndex[]: for each server, index of the next log entry to send to that server.
* matchIndex[]: for each server, index of highest log entry known to be replicated on server.

### AppendEntries RPC

#### Arguments

* term: leader's term
* leaderId: so follower can redirect clients
* prevLogIndex: index of log entry immediately preceding.
* prevLogTerm: term of preLogIndex entry
* entries[]: log entries to store
* leaderCommit: leader's commitIndex

#### Results

* term: currentTerm, for leader to update itself
* success: true if follower contained entry matching prevLogIndex and prevLogTerm

#### Reciever implementation

1. Reply false if term < currentTerm
2. Reply false if log does not contain an entry at pervLogIndex whose term matches prevLogTerm
3. If an existing entry conflicts with a new one (same index but different terms), delete the existing entry and all that follow it.
4. Append and new entries not already in the log
5. If leaderCommit > commitIndex, set commitIndex = min(leaderCommit, index of the new entry)

### RequestVote RPC

#### Arguments(Voting)

* term: candidate's term
* cnadidateId: cnadidate requesting vote
* lastLogIndex: index of candidate's last log entry 
* lastLogTerm: term of candidate's last log entry

#### Results(Voting)

* term: currentTerm, for candidate to udpdate itself
* voteGranted: true means candidate recieved vote

#### Reciever implementation(Voting)

1. Reply false if term < currentTerm
2. If votedFor is null or candidateId, and candidate's log is at least up-to-date as reciever's  log, grant vote.



