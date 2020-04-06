# Sirius
⚡️ Cloud Native Streaming Data RPC Proxy Based On gRPC

## Why Design this

More concise.To shield some details of accessing the data resource layer, the application only needs to pay attention to its own logic implementation.At the same time, you can avoid some resource-level problems, such as the number of resource connections.

Better monitoring. It is well known that monitoring at the data resource level often requires the maintenance of roles such as DBA. After abstracting the access to the data resource level, we can easily add the need for monitoring or distributed tracking.

Cloud native.Sirius allows front-end development and mobile development to have direct access to data resources, rather than through back-end servers. Can unlock more serverless computing potential.


## RoadMap

* Message Queue：
    * Kafka [In progress]

* Relational Database：
    * MySQL [TODO]

* NoSQL:
    * Redis [TODO]
    
## Other
The implementation of Sirius is very simple. It is a glue that connects gRPC and some drivers together. The meaning of this project is just to provide a new attempt. You are welcome to give suggestions for this project!


    
