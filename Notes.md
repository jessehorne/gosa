Notes on integrating with SimpleX Agents
===

I've decided to publish my notes to make it clear what I understand and what I do not for users of the library.

This document may save you some time or it might make things even more confusing.

# Commands

## New (INV)

### Client -> Agent

*Apparently you can only call this once without an error.*

```
[corrId]
[connId]
NEW T INV subscribe
```

### Agent -> Client (when other party JOINs)

```
[connId]
CONF [ID] [SMPServerAddress] [5DigitNumberMaybePort?]
[JSONDATA]
```

#### Notes

* `[connId]` is still confusing to me. I don't know what it is exactly.
* `[ID]` is like a unique identifier for the conversation I think
* `[JSONDATA]`...see `./gosa/joined.go JoinedUser`

