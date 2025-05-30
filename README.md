RHSM service
============

RHSM service providing VarLink API

Build
-----

To build RHSM service run:

```
go build
```

Run and Test
------------

To run the RHSM service open terminal:

```
sudo ./rhsm-service
```

To experiment with VarLnk API open another terminal.

Introspect interface:

```
sudo varlinkctl introspect /run/com.redhat.rhsm
```

To run some methods from VarLink interface run:

```
sudo varlinkctl call /run/com.redhat.rhsm com.redhat.rhsm.consumer.GetUUID '{"locale": "en_US.UTF-8"}'
```

```
sudo varlinkctl call /run/com.redhat.rhsm com.redhat.rhsm.consumer.GetOrg '{"locale": "en_US.UTF-8"}'
```
